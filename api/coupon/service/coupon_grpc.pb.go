// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/coupon/service/coupon.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CouponClient is the client API for Coupon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CouponClient interface {
	GetCouponByCategory(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Coupons, error)
	GetCouponByType(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*Coupons, error)
	GetUserCouponByStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Coupons, error)
	GetUserCouponByStatusWithCategory(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Coupons, error)
	CreateUserCoupon(ctx context.Context, in *CreateUserCouponRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type couponClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponClient(cc grpc.ClientConnInterface) CouponClient {
	return &couponClient{cc}
}

func (c *couponClient) GetCouponByCategory(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Coupons, error) {
	out := new(Coupons)
	err := c.cc.Invoke(ctx, "/coupon.service.Coupon/GetCouponByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponClient) GetCouponByType(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*Coupons, error) {
	out := new(Coupons)
	err := c.cc.Invoke(ctx, "/coupon.service.Coupon/GetCouponByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponClient) GetUserCouponByStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Coupons, error) {
	out := new(Coupons)
	err := c.cc.Invoke(ctx, "/coupon.service.Coupon/GetUserCouponByStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponClient) GetUserCouponByStatusWithCategory(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*Coupons, error) {
	out := new(Coupons)
	err := c.cc.Invoke(ctx, "/coupon.service.Coupon/GetUserCouponByStatusWithCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponClient) CreateUserCoupon(ctx context.Context, in *CreateUserCouponRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/coupon.service.Coupon/CreateUserCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponServer is the server API for Coupon service.
// All implementations must embed UnimplementedCouponServer
// for forward compatibility
type CouponServer interface {
	GetCouponByCategory(context.Context, *IdRequest) (*Coupons, error)
	GetCouponByType(context.Context, *TypeRequest) (*Coupons, error)
	GetUserCouponByStatus(context.Context, *StatusRequest) (*Coupons, error)
	GetUserCouponByStatusWithCategory(context.Context, *StatusRequest) (*Coupons, error)
	CreateUserCoupon(context.Context, *CreateUserCouponRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCouponServer()
}

// UnimplementedCouponServer must be embedded to have forward compatible implementations.
type UnimplementedCouponServer struct {
}

func (UnimplementedCouponServer) GetCouponByCategory(context.Context, *IdRequest) (*Coupons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponByCategory not implemented")
}
func (UnimplementedCouponServer) GetCouponByType(context.Context, *TypeRequest) (*Coupons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouponByType not implemented")
}
func (UnimplementedCouponServer) GetUserCouponByStatus(context.Context, *StatusRequest) (*Coupons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCouponByStatus not implemented")
}
func (UnimplementedCouponServer) GetUserCouponByStatusWithCategory(context.Context, *StatusRequest) (*Coupons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCouponByStatusWithCategory not implemented")
}
func (UnimplementedCouponServer) CreateUserCoupon(context.Context, *CreateUserCouponRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserCoupon not implemented")
}
func (UnimplementedCouponServer) mustEmbedUnimplementedCouponServer() {}

// UnsafeCouponServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponServer will
// result in compilation errors.
type UnsafeCouponServer interface {
	mustEmbedUnimplementedCouponServer()
}

func RegisterCouponServer(s grpc.ServiceRegistrar, srv CouponServer) {
	s.RegisterService(&Coupon_ServiceDesc, srv)
}

func _Coupon_GetCouponByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).GetCouponByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.service.Coupon/GetCouponByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).GetCouponByCategory(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coupon_GetCouponByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).GetCouponByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.service.Coupon/GetCouponByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).GetCouponByType(ctx, req.(*TypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coupon_GetUserCouponByStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).GetUserCouponByStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.service.Coupon/GetUserCouponByStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).GetUserCouponByStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coupon_GetUserCouponByStatusWithCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).GetUserCouponByStatusWithCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.service.Coupon/GetUserCouponByStatusWithCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).GetUserCouponByStatusWithCategory(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coupon_CreateUserCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).CreateUserCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.service.Coupon/CreateUserCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).CreateUserCoupon(ctx, req.(*CreateUserCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Coupon_ServiceDesc is the grpc.ServiceDesc for Coupon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coupon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coupon.service.Coupon",
	HandlerType: (*CouponServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCouponByCategory",
			Handler:    _Coupon_GetCouponByCategory_Handler,
		},
		{
			MethodName: "GetCouponByType",
			Handler:    _Coupon_GetCouponByType_Handler,
		},
		{
			MethodName: "GetUserCouponByStatus",
			Handler:    _Coupon_GetUserCouponByStatus_Handler,
		},
		{
			MethodName: "GetUserCouponByStatusWithCategory",
			Handler:    _Coupon_GetUserCouponByStatusWithCategory_Handler,
		},
		{
			MethodName: "CreateUserCoupon",
			Handler:    _Coupon_CreateUserCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/coupon/service/coupon.proto",
}
