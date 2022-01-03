// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HttpApiClient is the client API for HttpApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpApiClient interface {
	Echo(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserMessage, error)
}

type httpApiClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpApiClient(cc grpc.ClientConnInterface) HttpApiClient {
	return &httpApiClient{cc}
}

func (c *httpApiClient) Echo(ctx context.Context, in *UserMessage, opts ...grpc.CallOption) (*UserMessage, error) {
	out := new(UserMessage)
	err := c.cc.Invoke(ctx, "/main.HttpApi/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpApiServer is the server API for HttpApi service.
// All implementations must embed UnimplementedHttpApiServer
// for forward compatibility
type HttpApiServer interface {
	Echo(context.Context, *UserMessage) (*UserMessage, error)
	mustEmbedUnimplementedHttpApiServer()
}

// UnimplementedHttpApiServer must be embedded to have forward compatible implementations.
type UnimplementedHttpApiServer struct {
}

func (UnimplementedHttpApiServer) Echo(context.Context, *UserMessage) (*UserMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedHttpApiServer) mustEmbedUnimplementedHttpApiServer() {}

// UnsafeHttpApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpApiServer will
// result in compilation errors.
type UnsafeHttpApiServer interface {
	mustEmbedUnimplementedHttpApiServer()
}

func RegisterHttpApiServer(s grpc.ServiceRegistrar, srv HttpApiServer) {
	s.RegisterService(&HttpApi_ServiceDesc, srv)
}

func _HttpApi_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpApiServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.HttpApi/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpApiServer).Echo(ctx, req.(*UserMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// HttpApi_ServiceDesc is the grpc.ServiceDesc for HttpApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.HttpApi",
	HandlerType: (*HttpApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _HttpApi_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customApi.proto",
}
