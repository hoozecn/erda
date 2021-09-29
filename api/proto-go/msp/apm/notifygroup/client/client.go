// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: notifygroup.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/notifygroup/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// NotifyGroupService notifygroup.proto
	NotifyGroupService() pb.NotifyGroupServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		notifyGroupService: pb.NewNotifyGroupServiceClient(cc),
	}
}

type serviceClients struct {
	notifyGroupService pb.NotifyGroupServiceClient
}

func (c *serviceClients) NotifyGroupService() pb.NotifyGroupServiceClient {
	return c.notifyGroupService
}

type notifyGroupServiceWrapper struct {
	client pb.NotifyGroupServiceClient
	opts   []grpc1.CallOption
}

func (s *notifyGroupServiceWrapper) CreateNotifyGroup(ctx context.Context, req *pb.CreateNotifyGroupRequest) (*pb.CreateNotifyGroupResponse, error) {
	return s.client.CreateNotifyGroup(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyGroupServiceWrapper) QueryNotifyGroup(ctx context.Context, req *pb.QueryNotifyGroupRequest) (*pb.QueryNotifyGroupResponse, error) {
	return s.client.QueryNotifyGroup(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyGroupServiceWrapper) GetNotifyGroup(ctx context.Context, req *pb.GetNotifyGroupRequest) (*pb.GetNotifyGroupResponse, error) {
	return s.client.GetNotifyGroup(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyGroupServiceWrapper) UpdateNotifyGroup(ctx context.Context, req *pb.UpdateNotifyGroupRequest) (*pb.UpdateNotifyGroupResponse, error) {
	return s.client.UpdateNotifyGroup(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyGroupServiceWrapper) GetNotifyGroupDetail(ctx context.Context, req *pb.GetNotifyGroupDetailRequest) (*pb.GetNotifyGroupDetailResponse, error) {
	return s.client.GetNotifyGroupDetail(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyGroupServiceWrapper) DeleteNotifyGroup(ctx context.Context, req *pb.DeleteNotifyGroupRequest) (*pb.DeleteNotifyGroupResponse, error) {
	return s.client.DeleteNotifyGroup(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}