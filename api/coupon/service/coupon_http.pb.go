// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package service

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type CouponHTTPServer interface {
	CreateUserCoupon(context.Context, *CreateUserCouponRequest) (*emptypb.Empty, error)
	GetCouponByCategory(context.Context, *IdRequest) (*Coupons, error)
	GetCouponByType(context.Context, *TypeRequest) (*Coupons, error)
	GetUserCouponByStatus(context.Context, *StatusRequest) (*Coupons, error)
	GetUserCouponByStatusWithCategory(context.Context, *StatusRequest) (*Coupons, error)
}

func RegisterCouponHTTPServer(s *http.Server, srv CouponHTTPServer) {
	r := s.Route("/")
	r.GET("/coupon/category/{id}", _Coupon_GetCouponByCategory0_HTTP_Handler(srv))
	r.GET("/coupon/type/{type}", _Coupon_GetCouponByType0_HTTP_Handler(srv))
	r.GET("/coupon/user/status/{status}", _Coupon_GetUserCouponByStatus0_HTTP_Handler(srv))
	r.GET("/coupon/user/status/{status}/with_category", _Coupon_GetUserCouponByStatusWithCategory0_HTTP_Handler(srv))
	r.POST("/coupon/user", _Coupon_CreateUserCoupon0_HTTP_Handler(srv))
}

func _Coupon_GetCouponByCategory0_HTTP_Handler(srv CouponHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/coupon.service.Coupon/GetCouponByCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCouponByCategory(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Coupons)
		return ctx.Result(200, reply)
	}
}

func _Coupon_GetCouponByType0_HTTP_Handler(srv CouponHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TypeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/coupon.service.Coupon/GetCouponByType")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCouponByType(ctx, req.(*TypeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Coupons)
		return ctx.Result(200, reply)
	}
}

func _Coupon_GetUserCouponByStatus0_HTTP_Handler(srv CouponHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/coupon.service.Coupon/GetUserCouponByStatus")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserCouponByStatus(ctx, req.(*StatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Coupons)
		return ctx.Result(200, reply)
	}
}

func _Coupon_GetUserCouponByStatusWithCategory0_HTTP_Handler(srv CouponHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/coupon.service.Coupon/GetUserCouponByStatusWithCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserCouponByStatusWithCategory(ctx, req.(*StatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Coupons)
		return ctx.Result(200, reply)
	}
}

func _Coupon_CreateUserCoupon0_HTTP_Handler(srv CouponHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserCouponRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/coupon.service.Coupon/CreateUserCoupon")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUserCoupon(ctx, req.(*CreateUserCouponRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type CouponHTTPClient interface {
	CreateUserCoupon(ctx context.Context, req *CreateUserCouponRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetCouponByCategory(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	GetCouponByType(ctx context.Context, req *TypeRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	GetUserCouponByStatus(ctx context.Context, req *StatusRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	GetUserCouponByStatusWithCategory(ctx context.Context, req *StatusRequest, opts ...http.CallOption) (rsp *Coupons, err error)
}

type CouponHTTPClientImpl struct {
	cc *http.Client
}

func NewCouponHTTPClient(client *http.Client) CouponHTTPClient {
	return &CouponHTTPClientImpl{client}
}

func (c *CouponHTTPClientImpl) CreateUserCoupon(ctx context.Context, in *CreateUserCouponRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/coupon/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/coupon.service.Coupon/CreateUserCoupon"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CouponHTTPClientImpl) GetCouponByCategory(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/category/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/coupon.service.Coupon/GetCouponByCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CouponHTTPClientImpl) GetCouponByType(ctx context.Context, in *TypeRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/type/{type}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/coupon.service.Coupon/GetCouponByType"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CouponHTTPClientImpl) GetUserCouponByStatus(ctx context.Context, in *StatusRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/user/status/{status}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/coupon.service.Coupon/GetUserCouponByStatus"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CouponHTTPClientImpl) GetUserCouponByStatusWithCategory(ctx context.Context, in *StatusRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/user/status/{status}/with_category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/coupon.service.Coupon/GetUserCouponByStatusWithCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
