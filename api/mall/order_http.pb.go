// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

package mall

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

type OrderHTTPServer interface {
	CancelOrder(context.Context, *CancelOrderRequest) (*emptypb.Empty, error)
	GetOrder(context.Context, *GetOrderRequest) (*PlaceOrderReply, error)
	ListUserOrder(context.Context, *ListUserOrdersRequest) (*ListUserOrdersReply, error)
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderReply, error)
}

func RegisterOrderHTTPServer(s *http.Server, srv OrderHTTPServer) {
	r := s.Route("/")
	r.POST("/order", _Order_PlaceOrder0_HTTP_Handler(srv))
	r.GET("/orders", _Order_ListUserOrder0_HTTP_Handler(srv))
	r.GET("/order/{id}", _Order_GetOrder0_HTTP_Handler(srv))
	r.PUT("/order/cancel/{id}", _Order_CancelOrder0_HTTP_Handler(srv))
}

func _Order_PlaceOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PlaceOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.Order/PlaceOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PlaceOrder(ctx, req.(*PlaceOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlaceOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_ListUserOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUserOrdersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.Order/ListUserOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUserOrder(ctx, req.(*ListUserOrdersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUserOrdersReply)
		return ctx.Result(200, reply)
	}
}

func _Order_GetOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOrderRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.Order/GetOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrder(ctx, req.(*GetOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PlaceOrderReply)
		return ctx.Result(200, reply)
	}
}

func _Order_CancelOrder0_HTTP_Handler(srv OrderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CancelOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.Order/CancelOrder")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CancelOrder(ctx, req.(*CancelOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type OrderHTTPClient interface {
	CancelOrder(ctx context.Context, req *CancelOrderRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetOrder(ctx context.Context, req *GetOrderRequest, opts ...http.CallOption) (rsp *PlaceOrderReply, err error)
	ListUserOrder(ctx context.Context, req *ListUserOrdersRequest, opts ...http.CallOption) (rsp *ListUserOrdersReply, err error)
	PlaceOrder(ctx context.Context, req *PlaceOrderRequest, opts ...http.CallOption) (rsp *PlaceOrderReply, err error)
}

type OrderHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderHTTPClient(client *http.Client) OrderHTTPClient {
	return &OrderHTTPClientImpl{client}
}

func (c *OrderHTTPClientImpl) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/order/cancel/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/mall.Order/CancelOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...http.CallOption) (*PlaceOrderReply, error) {
	var out PlaceOrderReply
	pattern := "/order/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.Order/GetOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) ListUserOrder(ctx context.Context, in *ListUserOrdersRequest, opts ...http.CallOption) (*ListUserOrdersReply, error) {
	var out ListUserOrdersReply
	pattern := "/orders"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.Order/ListUserOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderHTTPClientImpl) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...http.CallOption) (*PlaceOrderReply, error) {
	var out PlaceOrderReply
	pattern := "/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/mall.Order/PlaceOrder"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
