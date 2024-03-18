// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: aggregator.proto

package _interface

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

const (
	AggregatorService_GetSmtMerkleProof_FullMethodName = "/aggregator.AggregatorService/GetSmtMerkleProof"
)

// AggregatorServiceClient is the client API for AggregatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatorServiceClient interface {
	GetSmtMerkleProof(ctx context.Context, in *SmtMerkleProofRequest, opts ...grpc.CallOption) (*SmtMerkleProofResponse, error)
}

type aggregatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatorServiceClient(cc grpc.ClientConnInterface) AggregatorServiceClient {
	return &aggregatorServiceClient{cc}
}

func (c *aggregatorServiceClient) GetSmtMerkleProof(ctx context.Context, in *SmtMerkleProofRequest, opts ...grpc.CallOption) (*SmtMerkleProofResponse, error) {
	out := new(SmtMerkleProofResponse)
	err := c.cc.Invoke(ctx, AggregatorService_GetSmtMerkleProof_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServiceServer is the server API for AggregatorService service.
// All implementations must embed UnimplementedAggregatorServiceServer
// for forward compatibility
type AggregatorServiceServer interface {
	GetSmtMerkleProof(context.Context, *SmtMerkleProofRequest) (*SmtMerkleProofResponse, error)
	mustEmbedUnimplementedAggregatorServiceServer()
}

// UnimplementedAggregatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAggregatorServiceServer struct {
}

func (UnimplementedAggregatorServiceServer) GetSmtMerkleProof(context.Context, *SmtMerkleProofRequest) (*SmtMerkleProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSmtMerkleProof not implemented")
}
func (UnimplementedAggregatorServiceServer) mustEmbedUnimplementedAggregatorServiceServer() {}

// UnsafeAggregatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatorServiceServer will
// result in compilation errors.
type UnsafeAggregatorServiceServer interface {
	mustEmbedUnimplementedAggregatorServiceServer()
}

func RegisterAggregatorServiceServer(s grpc.ServiceRegistrar, srv AggregatorServiceServer) {
	s.RegisterService(&AggregatorService_ServiceDesc, srv)
}

func _AggregatorService_GetSmtMerkleProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmtMerkleProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServiceServer).GetSmtMerkleProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregatorService_GetSmtMerkleProof_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServiceServer).GetSmtMerkleProof(ctx, req.(*SmtMerkleProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AggregatorService_ServiceDesc is the grpc.ServiceDesc for AggregatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aggregator.AggregatorService",
	HandlerType: (*AggregatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSmtMerkleProof",
			Handler:    _AggregatorService_GetSmtMerkleProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aggregator.proto",
}