// Copyright (c) MpesaOverlay. All rights reserved.
// Use of this source code is governed by a Apache-2.0 license that can be
// found in the LICENSE file.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: grpc/overlay.proto

package grpc

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
	Service_Token_FullMethodName             = "/mpesaoverlay.grpc.Service/Token"
	Service_ExpressQuery_FullMethodName      = "/mpesaoverlay.grpc.Service/ExpressQuery"
	Service_ExpressSimulate_FullMethodName   = "/mpesaoverlay.grpc.Service/ExpressSimulate"
	Service_B2CPayment_FullMethodName        = "/mpesaoverlay.grpc.Service/B2CPayment"
	Service_AccountBalance_FullMethodName    = "/mpesaoverlay.grpc.Service/AccountBalance"
	Service_C2BRegisterURL_FullMethodName    = "/mpesaoverlay.grpc.Service/C2BRegisterURL"
	Service_C2BSimulate_FullMethodName       = "/mpesaoverlay.grpc.Service/C2BSimulate"
	Service_GenerateQR_FullMethodName        = "/mpesaoverlay.grpc.Service/GenerateQR"
	Service_Reverse_FullMethodName           = "/mpesaoverlay.grpc.Service/Reverse"
	Service_TransactionStatus_FullMethodName = "/mpesaoverlay.grpc.Service/TransactionStatus"
	Service_RemitTax_FullMethodName          = "/mpesaoverlay.grpc.Service/RemitTax"
	Service_BusinessPayBill_FullMethodName   = "/mpesaoverlay.grpc.Service/BusinessPayBill"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Token(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TokenResp, error)
	ExpressQuery(ctx context.Context, in *ExpressQueryReq, opts ...grpc.CallOption) (*ExpressQueryResp, error)
	ExpressSimulate(ctx context.Context, in *ExpressSimulateReq, opts ...grpc.CallOption) (*ExpressSimulateResp, error)
	B2CPayment(ctx context.Context, in *B2CPaymentReq, opts ...grpc.CallOption) (*B2CPaymentResp, error)
	AccountBalance(ctx context.Context, in *AccountBalanceReq, opts ...grpc.CallOption) (*AccountBalanceResp, error)
	C2BRegisterURL(ctx context.Context, in *C2BRegisterURLReq, opts ...grpc.CallOption) (*C2BRegisterURLResp, error)
	C2BSimulate(ctx context.Context, in *C2BSimulateReq, opts ...grpc.CallOption) (*C2BSimulateResp, error)
	GenerateQR(ctx context.Context, in *GenerateQRReq, opts ...grpc.CallOption) (*GenerateQRResp, error)
	Reverse(ctx context.Context, in *ReverseReq, opts ...grpc.CallOption) (*ReverseResp, error)
	TransactionStatus(ctx context.Context, in *TransactionStatusReq, opts ...grpc.CallOption) (*TransactionStatusResp, error)
	RemitTax(ctx context.Context, in *RemitTaxReq, opts ...grpc.CallOption) (*RemitTaxResp, error)
	BusinessPayBill(ctx context.Context, in *BusinessPayBillReq, opts ...grpc.CallOption) (*BusinessPayBillResp, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Token(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TokenResp, error) {
	out := new(TokenResp)
	err := c.cc.Invoke(ctx, Service_Token_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ExpressQuery(ctx context.Context, in *ExpressQueryReq, opts ...grpc.CallOption) (*ExpressQueryResp, error) {
	out := new(ExpressQueryResp)
	err := c.cc.Invoke(ctx, Service_ExpressQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ExpressSimulate(ctx context.Context, in *ExpressSimulateReq, opts ...grpc.CallOption) (*ExpressSimulateResp, error) {
	out := new(ExpressSimulateResp)
	err := c.cc.Invoke(ctx, Service_ExpressSimulate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) B2CPayment(ctx context.Context, in *B2CPaymentReq, opts ...grpc.CallOption) (*B2CPaymentResp, error) {
	out := new(B2CPaymentResp)
	err := c.cc.Invoke(ctx, Service_B2CPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AccountBalance(ctx context.Context, in *AccountBalanceReq, opts ...grpc.CallOption) (*AccountBalanceResp, error) {
	out := new(AccountBalanceResp)
	err := c.cc.Invoke(ctx, Service_AccountBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) C2BRegisterURL(ctx context.Context, in *C2BRegisterURLReq, opts ...grpc.CallOption) (*C2BRegisterURLResp, error) {
	out := new(C2BRegisterURLResp)
	err := c.cc.Invoke(ctx, Service_C2BRegisterURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) C2BSimulate(ctx context.Context, in *C2BSimulateReq, opts ...grpc.CallOption) (*C2BSimulateResp, error) {
	out := new(C2BSimulateResp)
	err := c.cc.Invoke(ctx, Service_C2BSimulate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GenerateQR(ctx context.Context, in *GenerateQRReq, opts ...grpc.CallOption) (*GenerateQRResp, error) {
	out := new(GenerateQRResp)
	err := c.cc.Invoke(ctx, Service_GenerateQR_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Reverse(ctx context.Context, in *ReverseReq, opts ...grpc.CallOption) (*ReverseResp, error) {
	out := new(ReverseResp)
	err := c.cc.Invoke(ctx, Service_Reverse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TransactionStatus(ctx context.Context, in *TransactionStatusReq, opts ...grpc.CallOption) (*TransactionStatusResp, error) {
	out := new(TransactionStatusResp)
	err := c.cc.Invoke(ctx, Service_TransactionStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RemitTax(ctx context.Context, in *RemitTaxReq, opts ...grpc.CallOption) (*RemitTaxResp, error) {
	out := new(RemitTaxResp)
	err := c.cc.Invoke(ctx, Service_RemitTax_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BusinessPayBill(ctx context.Context, in *BusinessPayBillReq, opts ...grpc.CallOption) (*BusinessPayBillResp, error) {
	out := new(BusinessPayBillResp)
	err := c.cc.Invoke(ctx, Service_BusinessPayBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Token(context.Context, *Empty) (*TokenResp, error)
	ExpressQuery(context.Context, *ExpressQueryReq) (*ExpressQueryResp, error)
	ExpressSimulate(context.Context, *ExpressSimulateReq) (*ExpressSimulateResp, error)
	B2CPayment(context.Context, *B2CPaymentReq) (*B2CPaymentResp, error)
	AccountBalance(context.Context, *AccountBalanceReq) (*AccountBalanceResp, error)
	C2BRegisterURL(context.Context, *C2BRegisterURLReq) (*C2BRegisterURLResp, error)
	C2BSimulate(context.Context, *C2BSimulateReq) (*C2BSimulateResp, error)
	GenerateQR(context.Context, *GenerateQRReq) (*GenerateQRResp, error)
	Reverse(context.Context, *ReverseReq) (*ReverseResp, error)
	TransactionStatus(context.Context, *TransactionStatusReq) (*TransactionStatusResp, error)
	RemitTax(context.Context, *RemitTaxReq) (*RemitTaxResp, error)
	BusinessPayBill(context.Context, *BusinessPayBillReq) (*BusinessPayBillResp, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Token(context.Context, *Empty) (*TokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedServiceServer) ExpressQuery(context.Context, *ExpressQueryReq) (*ExpressQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpressQuery not implemented")
}
func (UnimplementedServiceServer) ExpressSimulate(context.Context, *ExpressSimulateReq) (*ExpressSimulateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpressSimulate not implemented")
}
func (UnimplementedServiceServer) B2CPayment(context.Context, *B2CPaymentReq) (*B2CPaymentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method B2CPayment not implemented")
}
func (UnimplementedServiceServer) AccountBalance(context.Context, *AccountBalanceReq) (*AccountBalanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountBalance not implemented")
}
func (UnimplementedServiceServer) C2BRegisterURL(context.Context, *C2BRegisterURLReq) (*C2BRegisterURLResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method C2BRegisterURL not implemented")
}
func (UnimplementedServiceServer) C2BSimulate(context.Context, *C2BSimulateReq) (*C2BSimulateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method C2BSimulate not implemented")
}
func (UnimplementedServiceServer) GenerateQR(context.Context, *GenerateQRReq) (*GenerateQRResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateQR not implemented")
}
func (UnimplementedServiceServer) Reverse(context.Context, *ReverseReq) (*ReverseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reverse not implemented")
}
func (UnimplementedServiceServer) TransactionStatus(context.Context, *TransactionStatusReq) (*TransactionStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionStatus not implemented")
}
func (UnimplementedServiceServer) RemitTax(context.Context, *RemitTaxReq) (*RemitTaxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemitTax not implemented")
}
func (UnimplementedServiceServer) BusinessPayBill(context.Context, *BusinessPayBillReq) (*BusinessPayBillResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BusinessPayBill not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Token_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Token(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ExpressQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ExpressQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ExpressQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ExpressQuery(ctx, req.(*ExpressQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ExpressSimulate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressSimulateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ExpressSimulate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_ExpressSimulate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ExpressSimulate(ctx, req.(*ExpressSimulateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_B2CPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(B2CPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).B2CPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_B2CPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).B2CPayment(ctx, req.(*B2CPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AccountBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AccountBalance(ctx, req.(*AccountBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_C2BRegisterURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2BRegisterURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).C2BRegisterURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_C2BRegisterURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).C2BRegisterURL(ctx, req.(*C2BRegisterURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_C2BSimulate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2BSimulateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).C2BSimulate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_C2BSimulate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).C2BSimulate(ctx, req.(*C2BSimulateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GenerateQR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateQRReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GenerateQR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GenerateQR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GenerateQR(ctx, req.(*GenerateQRReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Reverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Reverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_Reverse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Reverse(ctx, req.(*ReverseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TransactionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TransactionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_TransactionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TransactionStatus(ctx, req.(*TransactionStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RemitTax_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemitTaxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RemitTax(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_RemitTax_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RemitTax(ctx, req.(*RemitTaxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_BusinessPayBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BusinessPayBillReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).BusinessPayBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_BusinessPayBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).BusinessPayBill(ctx, req.(*BusinessPayBillReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mpesaoverlay.grpc.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Token",
			Handler:    _Service_Token_Handler,
		},
		{
			MethodName: "ExpressQuery",
			Handler:    _Service_ExpressQuery_Handler,
		},
		{
			MethodName: "ExpressSimulate",
			Handler:    _Service_ExpressSimulate_Handler,
		},
		{
			MethodName: "B2CPayment",
			Handler:    _Service_B2CPayment_Handler,
		},
		{
			MethodName: "AccountBalance",
			Handler:    _Service_AccountBalance_Handler,
		},
		{
			MethodName: "C2BRegisterURL",
			Handler:    _Service_C2BRegisterURL_Handler,
		},
		{
			MethodName: "C2BSimulate",
			Handler:    _Service_C2BSimulate_Handler,
		},
		{
			MethodName: "GenerateQR",
			Handler:    _Service_GenerateQR_Handler,
		},
		{
			MethodName: "Reverse",
			Handler:    _Service_Reverse_Handler,
		},
		{
			MethodName: "TransactionStatus",
			Handler:    _Service_TransactionStatus_Handler,
		},
		{
			MethodName: "RemitTax",
			Handler:    _Service_RemitTax_Handler,
		},
		{
			MethodName: "BusinessPayBill",
			Handler:    _Service_BusinessPayBill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/overlay.proto",
}
