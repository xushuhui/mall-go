// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

package app

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

type AppHTTPServer interface {
	CreateUserCoupon(context.Context, *CreateUserCouponRequest) (*emptypb.Empty, error)
	GetActivityByName(context.Context, *NameRequest) (*Activity, error)
	GetActivityWithCoupon(context.Context, *NameRequest) (*ActivityCoupon, error)
	GetBannerById(context.Context, *IdRequest) (*Banner, error)
	GetBannerByName(context.Context, *NameRequest) (*Banner, error)
	GetCouponByCategory(context.Context, *IdRequest) (*Coupons, error)
	GetCouponByType(context.Context, *TypeRequest) (*Coupons, error)
	GetSaleExplain(context.Context, *emptypb.Empty) (*SaleExplains, error)
	GetSpuByCategory(context.Context, *IdRequest) (*SpuPage, error)
	GetSpuById(context.Context, *IdRequest) (*SpuDetail, error)
	GetSpuLatest(context.Context, *emptypb.Empty) (*SpuPage, error)
	GetTagByType(context.Context, *TypeRequest) (*Tags, error)
	GetThemeByName(context.Context, *NameRequest) (*Theme, error)
	GetThemeByNames(context.Context, *ThemeByNamesRequest) (*Themes, error)
	GetUserCouponByStatus(context.Context, *StatusRequest) (*Coupons, error)
	GetUserCouponByStatusWithCategory(context.Context, *StatusRequest) (*Coupons, error)
	ListCategory(context.Context, *emptypb.Empty) (*Categories, error)
	ListGridCategory(context.Context, *emptypb.Empty) (*GridCategories, error)
}

func RegisterAppHTTPServer(s *http.Server, srv AppHTTPServer) {
	r := s.Route("/")
	r.GET("/banner/id/{id}", _App_GetBannerById0_HTTP_Handler(srv))
	r.GET("/banner/name/{name}", _App_GetBannerByName0_HTTP_Handler(srv))
	r.GET("/theme/by/names", _App_GetThemeByNames0_HTTP_Handler(srv))
	r.GET("/theme/name/{name}", _App_GetThemeByName0_HTTP_Handler(srv))
	r.GET("/activity/name/{name}", _App_GetActivityByName0_HTTP_Handler(srv))
	r.GET("/activity/name/{name}/with_coupon", _App_GetActivityWithCoupon0_HTTP_Handler(srv))
	r.GET("/category", _App_ListCategory0_HTTP_Handler(srv))
	r.GET("/category/grid", _App_ListGridCategory0_HTTP_Handler(srv))
	r.GET("/tag/type/{type}", _App_GetTagByType0_HTTP_Handler(srv))
	r.GET("/coupon/category/{id}", _App_GetCouponByCategory0_HTTP_Handler(srv))
	r.GET("/coupon/type/{type}", _App_GetCouponByType0_HTTP_Handler(srv))
	r.GET("/coupon/user/status/{status}", _App_GetUserCouponByStatus0_HTTP_Handler(srv))
	r.GET("/coupon/user/status/{status}/with_category", _App_GetUserCouponByStatusWithCategory0_HTTP_Handler(srv))
	r.POST("/coupon/user", _App_CreateUserCoupon0_HTTP_Handler(srv))
	r.GET("/sale_explain", _App_GetSaleExplain0_HTTP_Handler(srv))
	r.GET("/spu/{id}", _App_GetSpuById0_HTTP_Handler(srv))
	r.GET("/spu/latest", _App_GetSpuLatest0_HTTP_Handler(srv))
	r.GET("/spu/category/{id}", _App_GetSpuByCategory0_HTTP_Handler(srv))
}

func _App_GetBannerById0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetBannerById")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBannerById(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Banner)
		return ctx.Result(200, reply)
	}
}

func _App_GetBannerByName0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetBannerByName")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBannerByName(ctx, req.(*NameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Banner)
		return ctx.Result(200, reply)
	}
}

func _App_GetThemeByNames0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ThemeByNamesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetThemeByNames")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetThemeByNames(ctx, req.(*ThemeByNamesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Themes)
		return ctx.Result(200, reply)
	}
}

func _App_GetThemeByName0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetThemeByName")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetThemeByName(ctx, req.(*NameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Theme)
		return ctx.Result(200, reply)
	}
}

func _App_GetActivityByName0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetActivityByName")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetActivityByName(ctx, req.(*NameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Activity)
		return ctx.Result(200, reply)
	}
}

func _App_GetActivityWithCoupon0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in NameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetActivityWithCoupon")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetActivityWithCoupon(ctx, req.(*NameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ActivityCoupon)
		return ctx.Result(200, reply)
	}
}

func _App_ListCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/ListCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Categories)
		return ctx.Result(200, reply)
	}
}

func _App_ListGridCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/ListGridCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGridCategory(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GridCategories)
		return ctx.Result(200, reply)
	}
}

func _App_GetTagByType0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TypeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetTagByType")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTagByType(ctx, req.(*TypeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Tags)
		return ctx.Result(200, reply)
	}
}

func _App_GetCouponByCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetCouponByCategory")
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

func _App_GetCouponByType0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TypeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetCouponByType")
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

func _App_GetUserCouponByStatus0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetUserCouponByStatus")
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

func _App_GetUserCouponByStatusWithCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetUserCouponByStatusWithCategory")
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

func _App_CreateUserCoupon0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserCouponRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/CreateUserCoupon")
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

func _App_GetSaleExplain0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetSaleExplain")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSaleExplain(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaleExplains)
		return ctx.Result(200, reply)
	}
}

func _App_GetSpuById0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetSpuById")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSpuById(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SpuDetail)
		return ctx.Result(200, reply)
	}
}

func _App_GetSpuLatest0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetSpuLatest")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSpuLatest(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SpuPage)
		return ctx.Result(200, reply)
	}
}

func _App_GetSpuByCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/app.App/GetSpuByCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSpuByCategory(ctx, req.(*IdRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SpuPage)
		return ctx.Result(200, reply)
	}
}

type AppHTTPClient interface {
	CreateUserCoupon(ctx context.Context, req *CreateUserCouponRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetActivityByName(ctx context.Context, req *NameRequest, opts ...http.CallOption) (rsp *Activity, err error)
	GetActivityWithCoupon(ctx context.Context, req *NameRequest, opts ...http.CallOption) (rsp *ActivityCoupon, err error)
	GetBannerById(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *Banner, err error)
	GetBannerByName(ctx context.Context, req *NameRequest, opts ...http.CallOption) (rsp *Banner, err error)
	GetCouponByCategory(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	GetCouponByType(ctx context.Context, req *TypeRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	GetSaleExplain(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *SaleExplains, err error)
	GetSpuByCategory(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *SpuPage, err error)
	GetSpuById(ctx context.Context, req *IdRequest, opts ...http.CallOption) (rsp *SpuDetail, err error)
	GetSpuLatest(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *SpuPage, err error)
	GetTagByType(ctx context.Context, req *TypeRequest, opts ...http.CallOption) (rsp *Tags, err error)
	GetThemeByName(ctx context.Context, req *NameRequest, opts ...http.CallOption) (rsp *Theme, err error)
	GetThemeByNames(ctx context.Context, req *ThemeByNamesRequest, opts ...http.CallOption) (rsp *Themes, err error)
	GetUserCouponByStatus(ctx context.Context, req *StatusRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	GetUserCouponByStatusWithCategory(ctx context.Context, req *StatusRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	ListCategory(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Categories, err error)
	ListGridCategory(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GridCategories, err error)
}

type AppHTTPClientImpl struct {
	cc *http.Client
}

func NewAppHTTPClient(client *http.Client) AppHTTPClient {
	return &AppHTTPClientImpl{client}
}

func (c *AppHTTPClientImpl) CreateUserCoupon(ctx context.Context, in *CreateUserCouponRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/coupon/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/app.App/CreateUserCoupon"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetActivityByName(ctx context.Context, in *NameRequest, opts ...http.CallOption) (*Activity, error) {
	var out Activity
	pattern := "/activity/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetActivityByName"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetActivityWithCoupon(ctx context.Context, in *NameRequest, opts ...http.CallOption) (*ActivityCoupon, error) {
	var out ActivityCoupon
	pattern := "/activity/name/{name}/with_coupon"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetActivityWithCoupon"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetBannerById(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*Banner, error) {
	var out Banner
	pattern := "/banner/id/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetBannerById"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetBannerByName(ctx context.Context, in *NameRequest, opts ...http.CallOption) (*Banner, error) {
	var out Banner
	pattern := "/banner/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetBannerByName"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetCouponByCategory(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/category/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetCouponByCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetCouponByType(ctx context.Context, in *TypeRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/type/{type}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetCouponByType"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetSaleExplain(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*SaleExplains, error) {
	var out SaleExplains
	pattern := "/sale_explain"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetSaleExplain"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetSpuByCategory(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*SpuPage, error) {
	var out SpuPage
	pattern := "/spu/category/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetSpuByCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetSpuById(ctx context.Context, in *IdRequest, opts ...http.CallOption) (*SpuDetail, error) {
	var out SpuDetail
	pattern := "/spu/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetSpuById"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetSpuLatest(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*SpuPage, error) {
	var out SpuPage
	pattern := "/spu/latest"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetSpuLatest"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetTagByType(ctx context.Context, in *TypeRequest, opts ...http.CallOption) (*Tags, error) {
	var out Tags
	pattern := "/tag/type/{type}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetTagByType"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetThemeByName(ctx context.Context, in *NameRequest, opts ...http.CallOption) (*Theme, error) {
	var out Theme
	pattern := "/theme/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetThemeByName"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetThemeByNames(ctx context.Context, in *ThemeByNamesRequest, opts ...http.CallOption) (*Themes, error) {
	var out Themes
	pattern := "/theme/by/names"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetThemeByNames"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetUserCouponByStatus(ctx context.Context, in *StatusRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/user/status/{status}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetUserCouponByStatus"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetUserCouponByStatusWithCategory(ctx context.Context, in *StatusRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/user/status/{status}/with_category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/GetUserCouponByStatusWithCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) ListCategory(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Categories, error) {
	var out Categories
	pattern := "/category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/ListCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) ListGridCategory(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GridCategories, error) {
	var out GridCategories
	pattern := "/category/grid"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/app.App/ListGridCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
