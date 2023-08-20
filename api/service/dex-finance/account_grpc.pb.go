// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service/dex-finance/account.proto

package dex_finance

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
	Account_GetAccount_FullMethodName             = "/api.service.dex_finance.Account/GetAccount"
	Account_AccountRechargeTry_FullMethodName     = "/api.service.dex_finance.Account/AccountRechargeTry"
	Account_AccountRechargeConfirm_FullMethodName = "/api.service.dex_finance.Account/AccountRechargeConfirm"
	Account_AccountRechargeCancel_FullMethodName  = "/api.service.dex_finance.Account/AccountRechargeCancel"
)

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error)
	AccountRechargeTry(ctx context.Context, in *AccountRechargeRequest, opts ...grpc.CallOption) (*AccountRechargeReply, error)
	AccountRechargeConfirm(ctx context.Context, in *AccountRechargeRequest, opts ...grpc.CallOption) (*AccountRechargeReply, error)
	AccountRechargeCancel(ctx context.Context, in *AccountRechargeRequest, opts ...grpc.CallOption) (*AccountRechargeReply, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error) {
	out := new(GetAccountReply)
	err := c.cc.Invoke(ctx, Account_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) AccountRechargeTry(ctx context.Context, in *AccountRechargeRequest, opts ...grpc.CallOption) (*AccountRechargeReply, error) {
	out := new(AccountRechargeReply)
	err := c.cc.Invoke(ctx, Account_AccountRechargeTry_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) AccountRechargeConfirm(ctx context.Context, in *AccountRechargeRequest, opts ...grpc.CallOption) (*AccountRechargeReply, error) {
	out := new(AccountRechargeReply)
	err := c.cc.Invoke(ctx, Account_AccountRechargeConfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) AccountRechargeCancel(ctx context.Context, in *AccountRechargeRequest, opts ...grpc.CallOption) (*AccountRechargeReply, error) {
	out := new(AccountRechargeReply)
	err := c.cc.Invoke(ctx, Account_AccountRechargeCancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error)
	AccountRechargeTry(context.Context, *AccountRechargeRequest) (*AccountRechargeReply, error)
	AccountRechargeConfirm(context.Context, *AccountRechargeRequest) (*AccountRechargeReply, error)
	AccountRechargeCancel(context.Context, *AccountRechargeRequest) (*AccountRechargeReply, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAccountServer) AccountRechargeTry(context.Context, *AccountRechargeRequest) (*AccountRechargeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountRechargeTry not implemented")
}
func (UnimplementedAccountServer) AccountRechargeConfirm(context.Context, *AccountRechargeRequest) (*AccountRechargeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountRechargeConfirm not implemented")
}
func (UnimplementedAccountServer) AccountRechargeCancel(context.Context, *AccountRechargeRequest) (*AccountRechargeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountRechargeCancel not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_AccountRechargeTry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRechargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).AccountRechargeTry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_AccountRechargeTry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).AccountRechargeTry(ctx, req.(*AccountRechargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_AccountRechargeConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRechargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).AccountRechargeConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_AccountRechargeConfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).AccountRechargeConfirm(ctx, req.(*AccountRechargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_AccountRechargeCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRechargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).AccountRechargeCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_AccountRechargeCancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).AccountRechargeCancel(ctx, req.(*AccountRechargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.service.dex_finance.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _Account_GetAccount_Handler,
		},
		{
			MethodName: "AccountRechargeTry",
			Handler:    _Account_AccountRechargeTry_Handler,
		},
		{
			MethodName: "AccountRechargeConfirm",
			Handler:    _Account_AccountRechargeConfirm_Handler,
		},
		{
			MethodName: "AccountRechargeCancel",
			Handler:    _Account_AccountRechargeCancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/dex-finance/account.proto",
}
