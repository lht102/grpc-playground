// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/adder/adder.proto

package adder

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

// AdderServiceClient is the client API for AdderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdderServiceClient interface {
	AddNumbers(ctx context.Context, in *AddNumbersRequest, opts ...grpc.CallOption) (*AddNumbersResponse, error)
}

type adderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderServiceClient(cc grpc.ClientConnInterface) AdderServiceClient {
	return &adderServiceClient{cc}
}

func (c *adderServiceClient) AddNumbers(ctx context.Context, in *AddNumbersRequest, opts ...grpc.CallOption) (*AddNumbersResponse, error) {
	out := new(AddNumbersResponse)
	err := c.cc.Invoke(ctx, "/adder.AdderService/AddNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServiceServer is the server API for AdderService service.
// All implementations must embed UnimplementedAdderServiceServer
// for forward compatibility
type AdderServiceServer interface {
	AddNumbers(context.Context, *AddNumbersRequest) (*AddNumbersResponse, error)
	mustEmbedUnimplementedAdderServiceServer()
}

// UnimplementedAdderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdderServiceServer struct {
}

func (UnimplementedAdderServiceServer) AddNumbers(context.Context, *AddNumbersRequest) (*AddNumbersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNumbers not implemented")
}
func (UnimplementedAdderServiceServer) mustEmbedUnimplementedAdderServiceServer() {}

// UnsafeAdderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdderServiceServer will
// result in compilation errors.
type UnsafeAdderServiceServer interface {
	mustEmbedUnimplementedAdderServiceServer()
}

func RegisterAdderServiceServer(s grpc.ServiceRegistrar, srv AdderServiceServer) {
	s.RegisterService(&AdderService_ServiceDesc, srv)
}

func _AdderService_AddNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNumbersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServiceServer).AddNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adder.AdderService/AddNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServiceServer).AddNumbers(ctx, req.(*AddNumbersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdderService_ServiceDesc is the grpc.ServiceDesc for AdderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "adder.AdderService",
	HandlerType: (*AdderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNumbers",
			Handler:    _AdderService_AddNumbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/adder/adder.proto",
}
