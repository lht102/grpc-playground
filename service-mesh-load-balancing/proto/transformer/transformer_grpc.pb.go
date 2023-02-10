// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/transformer/transformer.proto

package transformer

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

// TransformerServiceClient is the client API for TransformerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransformerServiceClient interface {
	TransformNumber(ctx context.Context, in *TransformNumberRequest, opts ...grpc.CallOption) (*TransformNumberResponse, error)
}

type transformerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransformerServiceClient(cc grpc.ClientConnInterface) TransformerServiceClient {
	return &transformerServiceClient{cc}
}

func (c *transformerServiceClient) TransformNumber(ctx context.Context, in *TransformNumberRequest, opts ...grpc.CallOption) (*TransformNumberResponse, error) {
	out := new(TransformNumberResponse)
	err := c.cc.Invoke(ctx, "/transformer.TransformerService/TransformNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransformerServiceServer is the server API for TransformerService service.
// All implementations must embed UnimplementedTransformerServiceServer
// for forward compatibility
type TransformerServiceServer interface {
	TransformNumber(context.Context, *TransformNumberRequest) (*TransformNumberResponse, error)
	mustEmbedUnimplementedTransformerServiceServer()
}

// UnimplementedTransformerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransformerServiceServer struct {
}

func (UnimplementedTransformerServiceServer) TransformNumber(context.Context, *TransformNumberRequest) (*TransformNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransformNumber not implemented")
}
func (UnimplementedTransformerServiceServer) mustEmbedUnimplementedTransformerServiceServer() {}

// UnsafeTransformerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransformerServiceServer will
// result in compilation errors.
type UnsafeTransformerServiceServer interface {
	mustEmbedUnimplementedTransformerServiceServer()
}

func RegisterTransformerServiceServer(s grpc.ServiceRegistrar, srv TransformerServiceServer) {
	s.RegisterService(&TransformerService_ServiceDesc, srv)
}

func _TransformerService_TransformNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransformNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServiceServer).TransformNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transformer.TransformerService/TransformNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServiceServer).TransformNumber(ctx, req.(*TransformNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransformerService_ServiceDesc is the grpc.ServiceDesc for TransformerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransformerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transformer.TransformerService",
	HandlerType: (*TransformerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransformNumber",
			Handler:    _TransformerService_TransformNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transformer/transformer.proto",
}