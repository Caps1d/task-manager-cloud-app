// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: notification/pb/notification.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	NotificationService_NotifyUser_FullMethodName = "/notification.NotificationService/NotifyUser"
	NotificationService_NotifyTeam_FullMethodName = "/notification.NotificationService/NotifyTeam"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	NotifyUser(ctx context.Context, in *NotifyUserRequest, opts ...grpc.CallOption) (*NotifyUserResponse, error)
	NotifyTeam(ctx context.Context, in *NotifyTeamRequest, opts ...grpc.CallOption) (*NotifyTeamResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) NotifyUser(ctx context.Context, in *NotifyUserRequest, opts ...grpc.CallOption) (*NotifyUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NotifyUserResponse)
	err := c.cc.Invoke(ctx, NotificationService_NotifyUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) NotifyTeam(ctx context.Context, in *NotifyTeamRequest, opts ...grpc.CallOption) (*NotifyTeamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NotifyTeamResponse)
	err := c.cc.Invoke(ctx, NotificationService_NotifyTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	NotifyUser(context.Context, *NotifyUserRequest) (*NotifyUserResponse, error)
	NotifyTeam(context.Context, *NotifyTeamRequest) (*NotifyTeamResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) NotifyUser(context.Context, *NotifyUserRequest) (*NotifyUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUser not implemented")
}
func (UnimplementedNotificationServiceServer) NotifyTeam(context.Context, *NotifyTeamRequest) (*NotifyTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyTeam not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_NotifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).NotifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_NotifyUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).NotifyUser(ctx, req.(*NotifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_NotifyTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).NotifyTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_NotifyTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).NotifyTeam(ctx, req.(*NotifyTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyUser",
			Handler:    _NotificationService_NotifyUser_Handler,
		},
		{
			MethodName: "NotifyTeam",
			Handler:    _NotificationService_NotifyTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification/pb/notification.proto",
}
