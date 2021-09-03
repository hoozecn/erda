// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package podsTable

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/cmp/component-protocol/types"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

func init() {
	base.InitProviderWithCreator("cmp-dashboard-workload-detail", "podsTable", func() servicehub.Provider {
		return &ComponentPodsTable{}
	})
}

func (p *ComponentPodsTable) Render(ctx context.Context, component *cptype.Component, _ cptype.Scenario,
	event cptype.ComponentEvent, _ *cptype.GlobalStateData) error {
	p.InitComponent(ctx)
	if err := p.GenComponentState(component); err != nil {
		return fmt.Errorf("failed to gen podsTable component state, %v", err)
	}

	switch event.Operation {
	case cptype.InitializeOperation:
		p.State.PageNo = 1
		p.State.PageSize = 20
	case "changePageSize", "changeSort":
		p.State.PageNo = 1
	}

	if err := p.DecodeURLQuery(); err != nil {
		return fmt.Errorf("failed to decode url query for podsTable component, %v", err)
	}
	if err := p.RenderTable(); err != nil {
		return fmt.Errorf("failed to render podsTable component, %v", err)
	}
	if err := p.EncodeURLQuery(); err != nil {
		return fmt.Errorf("failed to encode url query for podsTable component, %v", err)
	}
	p.SetComponentValue()
	return nil
}

func (p *ComponentPodsTable) InitComponent(ctx context.Context) {
	bdl := ctx.Value(types.GlobalCtxKeyBundle).(*bundle.Bundle)
	p.bdl = bdl
	sdk := cputil.SDK(ctx)
	p.sdk = sdk
}

func (p *ComponentPodsTable) GenComponentState(component *cptype.Component) error {
	if component == nil || component.State == nil {
		return nil
	}
	var state State
	data, err := json.Marshal(component.State)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &state); err != nil {
		return err
	}
	p.State = state
	return nil
}

func (p *ComponentPodsTable) DecodeURLQuery() error {
	queryData, ok := p.sdk.InParams["workloadTable__urlQuery"].(string)
	if !ok {
		return nil
	}
	decode, err := base64.StdEncoding.DecodeString(queryData)
	if err != nil {
		return err
	}
	query := make(map[string]interface{})
	if err := json.Unmarshal(decode, &query); err != nil {
		return err
	}
	p.State.PageNo = int(query["pageNo"].(float64))
	p.State.PageSize = int(query["pageSize"].(float64))
	sorter := query["sorterData"].(map[string]interface{})
	p.State.Sorter.Field = sorter["field"].(string)
	p.State.Sorter.Order = sorter["order"].(string)
	return nil
}

func (p *ComponentPodsTable) EncodeURLQuery() error {
	query := make(map[string]interface{})
	query["pageNo"] = p.State.PageNo
	query["pageSize"] = p.State.PageSize
	query["sorterData"] = p.State.Sorter
	data, err := json.Marshal(query)
	if err != nil {
		return err
	}

	encode := base64.StdEncoding.EncodeToString(data)
	p.State.PodsTableURLQuery = encode
	return nil
}

func (p *ComponentPodsTable) RenderTable() error {
	userID := p.sdk.Identity.UserID
	orgID := p.sdk.Identity.OrgID
	workloadID := p.State.WorkloadID
	splits := strings.Split(workloadID, "_")
	if len(splits) != 3 {
		return fmt.Errorf("invalid workload id %s", workloadID)
	}
	kind, namespace, name := splits[0], splits[1], splits[2]

	req := apistructs.SteveRequest{
		UserID:      userID,
		OrgID:       orgID,
		Type:        apistructs.K8SResType(kind),
		ClusterName: p.State.ClusterName,
		Name:        name,
		Namespace:   namespace,
	}

	obj, err := p.bdl.GetSteveResource(&req)
	if err != nil {
		return err
	}

	labelSelectors := obj.Map("spec", "selector", "matchLabels")

	podReq := apistructs.SteveRequest{
		UserID:      userID,
		OrgID:       orgID,
		Type:        apistructs.K8SPod,
		ClusterName: p.State.ClusterName,
		Namespace:   namespace,
	}

	obj, err = p.bdl.ListSteveResource(&podReq)
	if err != nil {
		return err
	}
	list := obj.Slice("data")
	var items []Item
	for _, obj := range list {
		labels := obj.Map("metadata", "labels")
		if !matchSelector(labelSelectors, labels) {
			continue
		}

		name := obj.String("metadata", "name")
		namespace := obj.String("metadata", "namespace")
		fields := obj.StringSlice("metadata", "fields")
		if len(fields) != 9 {
			logrus.Errorf("length of pod %s:%s fields is invalid", namespace, name)
			continue
		}
		status := parsePodStatus(fields[2])

		containers := obj.Slice("spec", "containers")
		cpuRequests := resource.NewQuantity(0, resource.DecimalSI)
		cpuLimits := resource.NewQuantity(0, resource.DecimalSI)
		memRequests := resource.NewQuantity(0, resource.BinarySI)
		memLimits := resource.NewQuantity(0, resource.BinarySI)
		for _, container := range containers {
			cpuRequests.Add(*parseResource(container.String("resources", "requests", "cpu"), resource.DecimalSI))
			cpuLimits.Add(*parseResource(container.String("resources", "limits", "cpu"), resource.DecimalSI))
			memRequests.Add(*parseResource(container.String("resources", "requests", "memory"), resource.BinarySI))
			memLimits.Add(*parseResource(container.String("resources", "limits", "memory"), resource.BinarySI))
		}

		items = append(items, Item{
			Status: status,
			Name: Link{
				RenderType: "linkText",
				Value:      name,
				Operations: map[string]interface{}{
					"click": LinkOperation{
						Command: Command{
							Key:    "goto",
							Target: "cmpClustersPodDetail",
							State: CommandState{
								Query: map[string]string{
									"podName":   name,
									"namespace": namespace,
								},
							},
							JumpOut: true,
						},
						Reload: false,
					},
				},
			},
			Namespace:   namespace,
			IP:          fields[5],
			CPURequests: cpuRequests.String(),
			CPUPercent: Percent{ // TODO
				RenderType: "progress",
				Value:      "0",
				Tip:        "0",
				Status:     "success",
			},
			CPULimits:      cpuLimits.String(),
			MemoryRequests: memRequests.String(),
			MemoryPercent: Percent{
				RenderType: "progress",
				Value:      "0",
				Tip:        "0",
				Status:     "success",
			},
			MemoryLimits: memLimits.String(),
			Ready:        fields[1],
		})
	}

	if p.State.Sorter.Field != "" {
		cmpWrapper := func(field, order string) func(int, int) bool {
			ascend := order == "ascend"
			switch field {
			case "status":
				return func(i int, j int) bool {
					less := items[i].Status.Value < items[j].Status.Value
					if ascend {
						return less
					}
					return !less
				}
			case "name":
				return func(i int, j int) bool {
					less := items[i].Name.Value < items[j].Name.Value
					if ascend {
						return less
					}
					return !less
				}
			case "namespace":
				return func(i int, j int) bool {
					less := items[i].Namespace < items[j].Namespace
					if ascend {
						return less
					}
					return !less
				}
			case "ip":
				return func(i int, j int) bool {
					less := items[i].IP < items[j].IP
					if ascend {
						return less
					}
					return !less
				}
			case "cpuRequests":
				return func(i int, j int) bool {
					cpuI, _ := resource.ParseQuantity(items[i].CPURequests)
					cpuJ, _ := resource.ParseQuantity(items[j].CPURequests)
					less := cpuI.Cmp(cpuJ) < 0
					if ascend {
						return less
					}
					return !less
				}
			case "cpuPercent":
				return func(i int, j int) bool {
					vI, _ := strconv.ParseFloat(items[i].CPUPercent.Value, 64)
					vJ, _ := strconv.ParseFloat(items[j].CPUPercent.Value, 64)
					less := vI < vJ
					if ascend {
						return less
					}
					return !less
				}
			case "cpuLimits":
				return func(i int, j int) bool {
					cpuI, _ := resource.ParseQuantity(items[i].CPULimits)
					cpuJ, _ := resource.ParseQuantity(items[j].CPULimits)
					less := cpuI.Cmp(cpuJ) < 0
					if ascend {
						return less
					}
					return !less
				}
			case "memoryRequests":
				return func(i int, j int) bool {
					memI, _ := resource.ParseQuantity(items[i].MemoryRequests)
					memJ, _ := resource.ParseQuantity(items[j].MemoryRequests)
					less := memI.Cmp(memJ) < 0
					if ascend {
						return less
					}
					return !less
				}
			case "memoryPercent":
				return func(i int, j int) bool {
					vI, _ := strconv.ParseFloat(items[i].MemoryPercent.Value, 64)
					vJ, _ := strconv.ParseFloat(items[j].MemoryPercent.Value, 64)
					less := vI < vJ
					if ascend {
						return less
					}
					return !less
				}
			case "memoryLimits":
				return func(i int, j int) bool {
					memI, _ := resource.ParseQuantity(items[i].MemoryLimits)
					memJ, _ := resource.ParseQuantity(items[j].MemoryLimits)
					less := memI.Cmp(memJ) < 0
					if ascend {
						return less
					}
					return !less
				}
			case "ready":
				return func(i int, j int) bool {
					splits := strings.Split(items[i].Ready, "/")
					readyI := splits[0]
					splits = strings.Split(items[j].Ready, "/")
					readyJ := splits[0]
					less := readyI < readyJ
					if ascend {
						return less
					}
					return !less
				}
			default:
				return func(i int, j int) bool {
					return false
				}
			}
		}
		sort.Slice(items, cmpWrapper(p.State.Sorter.Field, p.State.Sorter.Order))
	}

	l, r := getRange(len(items), p.State.PageNo, p.State.PageSize)
	p.Data.List = items[l:r]
	p.State.Total = len(items)
	return nil
}

func (p *ComponentPodsTable) SetComponentValue() {
	p.Props.PageSizeOptions = []string{
		"10", "20", "50", "100",
	}
	p.Props.Columns = []Column{
		{
			DataIndex: "status",
			Title:     "Status",
			Width:     120,
			Sorter:    true,
		},
		{
			DataIndex: "name",
			Title:     "Name",
			Width:     180,
			Sorter:    true,
		},
		{
			DataIndex: "namespace",
			Title:     "Namespace",
			Width:     180,
			Sorter:    true,
		},
		{
			DataIndex: "ip",
			Title:     "IP",
			Width:     120,
			Sorter:    true,
		},
		{
			DataIndex: "cpuRequests",
			Title:     "CPU Requests",
			Width:     120,
			Sorter:    true,
		},
		{
			DataIndex: "cpuLimits",
			Title:     "CPU Limits",
			Width:     120,
			Sorter:    true,
		},
		{
			DataIndex: "cpuPercent",
			Title:     "CPU Percent",
			Width:     120,
			Sorter:    true,
		},
		{
			DataIndex: "memoryRequests",
			Title:     "Memory Requests",
			Width:     120,
			Sorter:    true,
		},
		{
			DataIndex: "memoryLimits",
			Title:     "Memory Limits",
			Width:     120,
			Sorter:    true,
		},
		{
			DataIndex: "memoryPercent",
			Title:     "Memory Percent",
			Width:     120,
			Sorter:    true,
		},
		{
			DataIndex: "ready",
			Title:     "Ready",
			Width:     80,
			Sorter:    true,
		},
	}

	p.Operations = map[string]interface{}{
		"changePageNo": Operation{
			Key:    "changePageNo",
			Reload: true,
		},
		"changePageSize": Operation{
			Key:    "changePageSize",
			Reload: true,
		},
		"changeSort": Operation{
			Key:    "changeSort",
			Reload: true,
		},
	}
}

func matchSelector(selector, labels map[string]interface{}) bool {
	for key, v := range selector {
		value, ok := v.(string)
		if !ok {
			return false
		}
		labelV, ok := labels[key]
		if !ok {
			return false
		}
		labelValue, ok := labelV.(string)
		if !ok || labelValue != value {
			return false
		}
	}
	return true
}

func parsePodStatus(state string) Status {
	status := Status{
		RenderType: "text",
		Value:      state,
	}
	switch state {
	case "Completed":
		status.StyleConfig.Color = "steelBlue"
	case "ContainerCreating":
		status.StyleConfig.Color = "orange"
	case "CrashLoopBackOff":
		status.StyleConfig.Color = "red"
	case "Error":
		status.StyleConfig.Color = "maroon"
	case "Evicted":
		status.StyleConfig.Color = "darkgoldenrod"
	case "ImagePullBackOff":
		status.StyleConfig.Color = "darksalmon"
	case "Pending":
		status.StyleConfig.Color = "teal"
	case "Running":
		status.StyleConfig.Color = "lightgreen"
	case "Terminating":
		status.StyleConfig.Color = "brown"
	}
	return status
}

func parseResource(str string, format resource.Format) *resource.Quantity {
	if str == "" {
		return resource.NewQuantity(0, format)
	}
	res, _ := resource.ParseQuantity(str)
	return &res
}

func getRange(length, pageNo, pageSize int) (int, int) {
	l := (pageNo - 1) * pageSize
	if l >= length || l < 0 {
		l = 0
	}
	r := l + pageSize
	if r > length || r < 0 {
		r = length
	}
	return l, r
}