// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: miniblog/v1/miniblog.proto

package miniblog

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MiniBlog_ListUser_FullMethodName = "/v1.MiniBlog/ListUser"
)

// MiniBlogClient is the client API for MiniBlog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MiniBlog 定义了一个MiniBlog RPC 服务.
type MiniBlogClient interface {
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
}

type miniBlogClient struct {
	cc grpc.ClientConnInterface
}

func NewMiniBlogClient(cc grpc.ClientConnInterface) MiniBlogClient {
	return &miniBlogClient{cc}
}

func (c *miniBlogClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, MiniBlog_ListUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MiniBlogServer is the server API for MiniBlog service.
// All implementations must embed UnimplementedMiniBlogServer
// for forward compatibility.
//
// MiniBlog 定义了一个MiniBlog RPC 服务.
type MiniBlogServer interface {
	ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error)
	mustEmbedUnimplementedMiniBlogServer()
}

// UnimplementedMiniBlogServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMiniBlogServer struct{}

func (UnimplementedMiniBlogServer) ListUser(context.Context, *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedMiniBlogServer) mustEmbedUnimplementedMiniBlogServer() {}
func (UnimplementedMiniBlogServer) testEmbeddedByValue()                  {}

// UnsafeMiniBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MiniBlogServer will
// result in compilation errors.
type UnsafeMiniBlogServer interface {
	mustEmbedUnimplementedMiniBlogServer()
}

func RegisterMiniBlogServer(s grpc.ServiceRegistrar, srv MiniBlogServer) {
	// If the following call pancis, it indicates UnimplementedMiniBlogServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MiniBlog_ServiceDesc, srv)
}

func _MiniBlog_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MiniBlogServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MiniBlog_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MiniBlogServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MiniBlog_ServiceDesc is the grpc.ServiceDesc for MiniBlog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MiniBlog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MiniBlog",
	HandlerType: (*MiniBlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUser",
			Handler:    _MiniBlog_ListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "miniblog/v1/miniblog.proto",
}
