// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package payclient

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

// PayClient is the client API for Pay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayClient interface {
	AppUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*UnifiedOrderResp, error)
	H5UnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*H5UnifiedOrderResp, error)
	JsUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*UnifiedOrderResp, error)
	OrderQuery(ctx context.Context, in *OrderQueryReq, opts ...grpc.CallOption) (*OrderQueryResp, error)
}

type payClient struct {
	cc grpc.ClientConnInterface
}

func NewPayClient(cc grpc.ClientConnInterface) PayClient {
	return &payClient{cc}
}

func (c *payClient) AppUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*UnifiedOrderResp, error) {
	out := new(UnifiedOrderResp)
	err := c.cc.Invoke(ctx, "/payclient.Pay/AppUnifiedOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) H5UnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*H5UnifiedOrderResp, error) {
	out := new(H5UnifiedOrderResp)
	err := c.cc.Invoke(ctx, "/payclient.Pay/H5UnifiedOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) JsUnifiedOrder(ctx context.Context, in *UnifiedOrderReq, opts ...grpc.CallOption) (*UnifiedOrderResp, error) {
	out := new(UnifiedOrderResp)
	err := c.cc.Invoke(ctx, "/payclient.Pay/JsUnifiedOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) OrderQuery(ctx context.Context, in *OrderQueryReq, opts ...grpc.CallOption) (*OrderQueryResp, error) {
	out := new(OrderQueryResp)
	err := c.cc.Invoke(ctx, "/payclient.Pay/OrderQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServer is the server API for Pay service.
// All implementations must embed UnimplementedPayServer
// for forward compatibility
type PayServer interface {
	AppUnifiedOrder(context.Context, *UnifiedOrderReq) (*UnifiedOrderResp, error)
	H5UnifiedOrder(context.Context, *UnifiedOrderReq) (*H5UnifiedOrderResp, error)
	JsUnifiedOrder(context.Context, *UnifiedOrderReq) (*UnifiedOrderResp, error)
	OrderQuery(context.Context, *OrderQueryReq) (*OrderQueryResp, error)
	mustEmbedUnimplementedPayServer()
}

// UnimplementedPayServer must be embedded to have forward compatible implementations.
type UnimplementedPayServer struct {
}

func (UnimplementedPayServer) AppUnifiedOrder(context.Context, *UnifiedOrderReq) (*UnifiedOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppUnifiedOrder not implemented")
}
func (UnimplementedPayServer) H5UnifiedOrder(context.Context, *UnifiedOrderReq) (*H5UnifiedOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method H5UnifiedOrder not implemented")
}
func (UnimplementedPayServer) JsUnifiedOrder(context.Context, *UnifiedOrderReq) (*UnifiedOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JsUnifiedOrder not implemented")
}
func (UnimplementedPayServer) OrderQuery(context.Context, *OrderQueryReq) (*OrderQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderQuery not implemented")
}
func (UnimplementedPayServer) mustEmbedUnimplementedPayServer() {}

// UnsafePayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayServer will
// result in compilation errors.
type UnsafePayServer interface {
	mustEmbedUnimplementedPayServer()
}

func RegisterPayServer(s grpc.ServiceRegistrar, srv PayServer) {
	s.RegisterService(&Pay_ServiceDesc, srv)
}

func _Pay_AppUnifiedOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnifiedOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).AppUnifiedOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payclient.Pay/AppUnifiedOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).AppUnifiedOrder(ctx, req.(*UnifiedOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_H5UnifiedOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnifiedOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).H5UnifiedOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payclient.Pay/H5UnifiedOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).H5UnifiedOrder(ctx, req.(*UnifiedOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_JsUnifiedOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnifiedOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).JsUnifiedOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payclient.Pay/JsUnifiedOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).JsUnifiedOrder(ctx, req.(*UnifiedOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_OrderQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).OrderQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payclient.Pay/OrderQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).OrderQuery(ctx, req.(*OrderQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pay_ServiceDesc is the grpc.ServiceDesc for Pay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payclient.Pay",
	HandlerType: (*PayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppUnifiedOrder",
			Handler:    _Pay_AppUnifiedOrder_Handler,
		},
		{
			MethodName: "H5UnifiedOrder",
			Handler:    _Pay_H5UnifiedOrder_Handler,
		},
		{
			MethodName: "JsUnifiedOrder",
			Handler:    _Pay_JsUnifiedOrder_Handler,
		},
		{
			MethodName: "OrderQuery",
			Handler:    _Pay_OrderQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay.proto",
}
