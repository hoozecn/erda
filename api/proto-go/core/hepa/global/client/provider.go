// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: global.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/global/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.hepa.global",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType             = reflect.TypeOf((*Client)(nil)).Elem()
	globalServiceClientType = reflect.TypeOf((*pb.GlobalServiceClient)(nil)).Elem()
	globalServiceServerType = reflect.TypeOf((*pb.GlobalServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.hepa.global-client":
		return p.client
	case "erda.core.hepa.global.GlobalService":
		return &globalServiceWrapper{client: p.client.GlobalService(), opts: opts}
	case "erda.core.hepa.global.GlobalService.client":
		return p.client.GlobalService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case globalServiceClientType:
		return p.client.GlobalService()
	case globalServiceServerType:
		return &globalServiceWrapper{client: p.client.GlobalService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.global-client", &servicehub.Spec{
		Services: []string{
			"erda.core.hepa.global.GlobalService",
			"erda.core.hepa.global-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			globalServiceClientType,
			// server types
			globalServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}