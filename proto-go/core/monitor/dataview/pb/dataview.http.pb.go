// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: dataview.proto

package pb

import (
	context "context"
	http1 "net/http"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// DataViewServiceHandler is the server API for DataViewService service.
type DataViewServiceHandler interface {
	// GET /api/dashboard/system/blocks
	ListSystemViews(context.Context, *ListSystemViewsRequest) (*ListSystemViewsResponse, error)
	// GET /api/dashboard/system/blocks/{id}
	GetSystemView(context.Context, *GetSystemViewRequest) (*GetSystemViewResponse, error)
	// GET /api/dashboard/blocks
	ListCustomViews(context.Context, *ListCustomViewsRequest) (*ListCustomViewsResponse, error)
	// GET /api/dashboard/blocks/{id}
	GetCustomView(context.Context, *GetCustomViewRequest) (*GetCustomViewResponse, error)
	// POST /api/dashboard/blocks
	CreateCustomView(context.Context, *CreateCustomViewRequest) (*CreateCustomViewResponse, error)
	// PUT /api/dashboard/blocks/{id}
	UpdateCustomView(context.Context, *UpdateCustomViewRequest) (*UpdateCustomViewResponse, error)
	// DELETE /api/dashboard/blocks/{id}
	DeleteCustomView(context.Context, *DeleteCustomViewRequest) (*DeleteCustomViewResponse, error)
}

// RegisterDataViewServiceHandler register DataViewServiceHandler to http.Router.
func RegisterDataViewServiceHandler(r http.Router, srv DataViewServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		return func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
	}

	add_ListSystemViews := func(method, path string, fn func(context.Context, *ListSystemViewsRequest) (*ListSystemViewsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListSystemViewsRequest))
		}
		var ListSystemViews_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListSystemViews_info = transport.NewServiceInfo("erda.core.monitor.dataview.DataViewService", "ListSystemViews", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListSystemViews_info)
				}
				r = r.WithContext(ctx)
				var in ListSystemViewsRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["scopeId"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetSystemView := func(method, path string, fn func(context.Context, *GetSystemViewRequest) (*GetSystemViewResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetSystemViewRequest))
		}
		var GetSystemView_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetSystemView_info = transport.NewServiceInfo("erda.core.monitor.dataview.DataViewService", "GetSystemView", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetSystemView_info)
				}
				r = r.WithContext(ctx)
				var in GetSystemViewRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListCustomViews := func(method, path string, fn func(context.Context, *ListCustomViewsRequest) (*ListCustomViewsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*ListCustomViewsRequest))
		}
		var ListCustomViews_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListCustomViews_info = transport.NewServiceInfo("erda.core.monitor.dataview.DataViewService", "ListCustomViews", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListCustomViews_info)
				}
				r = r.WithContext(ctx)
				var in ListCustomViewsRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				params := r.URL.Query()
				if vals := params["scopeId"]; len(vals) > 0 {
					in.ScopeID = vals[0]
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetCustomView := func(method, path string, fn func(context.Context, *GetCustomViewRequest) (*GetCustomViewResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetCustomViewRequest))
		}
		var GetCustomView_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetCustomView_info = transport.NewServiceInfo("erda.core.monitor.dataview.DataViewService", "GetCustomView", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetCustomView_info)
				}
				r = r.WithContext(ctx)
				var in GetCustomViewRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateCustomView := func(method, path string, fn func(context.Context, *CreateCustomViewRequest) (*CreateCustomViewResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateCustomViewRequest))
		}
		var CreateCustomView_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateCustomView_info = transport.NewServiceInfo("erda.core.monitor.dataview.DataViewService", "CreateCustomView", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateCustomView_info)
				}
				r = r.WithContext(ctx)
				var in CreateCustomViewRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_UpdateCustomView := func(method, path string, fn func(context.Context, *UpdateCustomViewRequest) (*UpdateCustomViewResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateCustomViewRequest))
		}
		var UpdateCustomView_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateCustomView_info = transport.NewServiceInfo("erda.core.monitor.dataview.DataViewService", "UpdateCustomView", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateCustomView_info)
				}
				r = r.WithContext(ctx)
				var in UpdateCustomViewRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_DeleteCustomView := func(method, path string, fn func(context.Context, *DeleteCustomViewRequest) (*DeleteCustomViewResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteCustomViewRequest))
		}
		var DeleteCustomView_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteCustomView_info = transport.NewServiceInfo("erda.core.monitor.dataview.DataViewService", "DeleteCustomView", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteCustomView_info)
				}
				r = r.WithContext(ctx)
				var in DeleteCustomViewRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "id":
							in.Id = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_ListSystemViews("GET", "/api/dashboard/system/blocks", srv.ListSystemViews)
	add_GetSystemView("GET", "/api/dashboard/system/blocks/{id}", srv.GetSystemView)
	add_ListCustomViews("GET", "/api/dashboard/blocks", srv.ListCustomViews)
	add_GetCustomView("GET", "/api/dashboard/blocks/{id}", srv.GetCustomView)
	add_CreateCustomView("POST", "/api/dashboard/blocks", srv.CreateCustomView)
	add_UpdateCustomView("PUT", "/api/dashboard/blocks/{id}", srv.UpdateCustomView)
	add_DeleteCustomView("DELETE", "/api/dashboard/blocks/{id}", srv.DeleteCustomView)
}