// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: cms.proto

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

// CmsServiceHandler is the server API for CmsService service.
type CmsServiceHandler interface {
	// POST /api/pipelines/cms/ns
	CreateNs(context.Context, *CmsCreateNsRequest) (*CmsCreateNsResponse, error)
	// GET /api/pipelines/cms/ns
	ListCmsNs(context.Context, *CmsListNsRequest) (*CmsListNsResponse, error)
	// POST /api/pipelines/cms/ns/{ns}
	UpdateCmsNsConfigs(context.Context, *CmsNsConfigsUpdateRequest) (*CmsNsConfigsUpdateResponse, error)
	// DELETE /api/pipelines/cms/ns/{ns}
	DeleteCmsNsConfigs(context.Context, *CmsNsConfigsDeleteRequest) (*CmsNsConfigsDeleteResponse, error)
	// GET /api/pipelines/cms/ns/{ns}
	GetCmsNsConfigs(context.Context, *CmsNsConfigsGetRequest) (*CmsNsConfigsGetResponse, error)
}

// RegisterCmsServiceHandler register CmsServiceHandler to http.Router.
func RegisterCmsServiceHandler(r http.Router, srv CmsServiceHandler, opts ...http.HandleOption) {
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

	add_CreateNs := func(method, path string, fn func(context.Context, *CmsCreateNsRequest) (*CmsCreateNsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CmsCreateNsRequest))
		}
		var CreateNs_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateNs_info = transport.NewServiceInfo("erda.core.pipeline.cms.CmsService", "CreateNs", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateNs_info)
				}
				r = r.WithContext(ctx)
				var in CmsCreateNsRequest
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

	add_ListCmsNs := func(method, path string, fn func(context.Context, *CmsListNsRequest) (*CmsListNsResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CmsListNsRequest))
		}
		var ListCmsNs_info transport.ServiceInfo
		if h.Interceptor != nil {
			ListCmsNs_info = transport.NewServiceInfo("erda.core.pipeline.cms.CmsService", "ListCmsNs", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, ListCmsNs_info)
				}
				r = r.WithContext(ctx)
				var in CmsListNsRequest
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

	add_UpdateCmsNsConfigs := func(method, path string, fn func(context.Context, *CmsNsConfigsUpdateRequest) (*CmsNsConfigsUpdateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CmsNsConfigsUpdateRequest))
		}
		var UpdateCmsNsConfigs_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateCmsNsConfigs_info = transport.NewServiceInfo("erda.core.pipeline.cms.CmsService", "UpdateCmsNsConfigs", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateCmsNsConfigs_info)
				}
				r = r.WithContext(ctx)
				var in CmsNsConfigsUpdateRequest
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
						case "ns":
							in.Ns = val
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

	add_DeleteCmsNsConfigs := func(method, path string, fn func(context.Context, *CmsNsConfigsDeleteRequest) (*CmsNsConfigsDeleteResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CmsNsConfigsDeleteRequest))
		}
		var DeleteCmsNsConfigs_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteCmsNsConfigs_info = transport.NewServiceInfo("erda.core.pipeline.cms.CmsService", "DeleteCmsNsConfigs", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteCmsNsConfigs_info)
				}
				r = r.WithContext(ctx)
				var in CmsNsConfigsDeleteRequest
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
						case "ns":
							in.Ns = val
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

	add_GetCmsNsConfigs := func(method, path string, fn func(context.Context, *CmsNsConfigsGetRequest) (*CmsNsConfigsGetResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CmsNsConfigsGetRequest))
		}
		var GetCmsNsConfigs_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetCmsNsConfigs_info = transport.NewServiceInfo("erda.core.pipeline.cms.CmsService", "GetCmsNsConfigs", srv)
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
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetCmsNsConfigs_info)
				}
				r = r.WithContext(ctx)
				var in CmsNsConfigsGetRequest
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
						case "ns":
							in.Ns = val
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

	add_CreateNs("POST", "/api/pipelines/cms/ns", srv.CreateNs)
	add_ListCmsNs("GET", "/api/pipelines/cms/ns", srv.ListCmsNs)
	add_UpdateCmsNsConfigs("POST", "/api/pipelines/cms/ns/{ns}", srv.UpdateCmsNsConfigs)
	add_DeleteCmsNsConfigs("DELETE", "/api/pipelines/cms/ns/{ns}", srv.DeleteCmsNsConfigs)
	add_GetCmsNsConfigs("GET", "/api/pipelines/cms/ns/{ns}", srv.GetCmsNsConfigs)
}