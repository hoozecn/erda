// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: label.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterLabelServiceImp label.proto
func RegisterLabelServiceImp(regester transport.Register, srv LabelServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterLabelServiceHandler(regester, LabelServiceHandler(srv), _ops.HTTP...)
	RegisterLabelServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.pipeline.label.LabelService",
	)
}

var (
	labelServiceClientType  = reflect.TypeOf((*LabelServiceClient)(nil)).Elem()
	labelServiceServerType  = reflect.TypeOf((*LabelServiceServer)(nil)).Elem()
	labelServiceHandlerType = reflect.TypeOf((*LabelServiceHandler)(nil)).Elem()
)

// LabelServiceClientType .
func LabelServiceClientType() reflect.Type { return labelServiceClientType }

// LabelServiceServerType .
func LabelServiceServerType() reflect.Type { return labelServiceServerType }

// LabelServiceHandlerType .
func LabelServiceHandlerType() reflect.Type { return labelServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		labelServiceClientType,
		// server types
		labelServiceServerType,
		// handler types
		labelServiceHandlerType,
	}
}