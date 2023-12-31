// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service/data/chain.proto

package data

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
	Chain_SyncBlockInfo_FullMethodName            = "/api.service.data.Chain/SyncBlockInfo"
	Chain_SyncBlockTransactionInfo_FullMethodName = "/api.service.data.Chain/SyncBlockTransactionInfo"
)

// ChainClient is the client API for Chain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainClient interface {
	SyncBlockInfo(ctx context.Context, in *SyncBlockInfoRequest, opts ...grpc.CallOption) (*SyncBlockInfoReply, error)
	SyncBlockTransactionInfo(ctx context.Context, in *SyncBlockTransactionInfoRequest, opts ...grpc.CallOption) (*SyncBlockTransactionInfoReply, error)
}

type chainClient struct {
	cc grpc.ClientConnInterface
}

func NewChainClient(cc grpc.ClientConnInterface) ChainClient {
	return &chainClient{cc}
}

func (c *chainClient) SyncBlockInfo(ctx context.Context, in *SyncBlockInfoRequest, opts ...grpc.CallOption) (*SyncBlockInfoReply, error) {
	out := new(SyncBlockInfoReply)
	err := c.cc.Invoke(ctx, Chain_SyncBlockInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainClient) SyncBlockTransactionInfo(ctx context.Context, in *SyncBlockTransactionInfoRequest, opts ...grpc.CallOption) (*SyncBlockTransactionInfoReply, error) {
	out := new(SyncBlockTransactionInfoReply)
	err := c.cc.Invoke(ctx, Chain_SyncBlockTransactionInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainServer is the server API for Chain service.
// All implementations must embed UnimplementedChainServer
// for forward compatibility
type ChainServer interface {
	SyncBlockInfo(context.Context, *SyncBlockInfoRequest) (*SyncBlockInfoReply, error)
	SyncBlockTransactionInfo(context.Context, *SyncBlockTransactionInfoRequest) (*SyncBlockTransactionInfoReply, error)
	mustEmbedUnimplementedChainServer()
}

// UnimplementedChainServer must be embedded to have forward compatible implementations.
type UnimplementedChainServer struct {
}

func (UnimplementedChainServer) SyncBlockInfo(context.Context, *SyncBlockInfoRequest) (*SyncBlockInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncBlockInfo not implemented")
}
func (UnimplementedChainServer) SyncBlockTransactionInfo(context.Context, *SyncBlockTransactionInfoRequest) (*SyncBlockTransactionInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncBlockTransactionInfo not implemented")
}
func (UnimplementedChainServer) mustEmbedUnimplementedChainServer() {}

// UnsafeChainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainServer will
// result in compilation errors.
type UnsafeChainServer interface {
	mustEmbedUnimplementedChainServer()
}

func RegisterChainServer(s grpc.ServiceRegistrar, srv ChainServer) {
	s.RegisterService(&Chain_ServiceDesc, srv)
}

func _Chain_SyncBlockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncBlockInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).SyncBlockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chain_SyncBlockInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).SyncBlockInfo(ctx, req.(*SyncBlockInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chain_SyncBlockTransactionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncBlockTransactionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).SyncBlockTransactionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chain_SyncBlockTransactionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).SyncBlockTransactionInfo(ctx, req.(*SyncBlockTransactionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chain_ServiceDesc is the grpc.ServiceDesc for Chain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.service.data.Chain",
	HandlerType: (*ChainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncBlockInfo",
			Handler:    _Chain_SyncBlockInfo_Handler,
		},
		{
			MethodName: "SyncBlockTransactionInfo",
			Handler:    _Chain_SyncBlockTransactionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/data/chain.proto",
}
