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

type AppHTTPServer interface {
	CollectCoupon(context.Context, *CollectCouponRequest) (*emptypb.Empty, error)
	GetActivityByName(context.Context, *ActivityByNameRequest) (*Activity, error)
	GetActivityWithCoupon(context.Context, *ActivityWithCouponRequest) (*ActivityCoupon, error)
	GetAllCategory(context.Context, *emptypb.Empty) (*AllCategory, error)
	GetBannerById(context.Context, *BannerByIdRequest) (*Banner, error)
	GetBannerByName(context.Context, *BannerByNameRequest) (*Banner, error)
	GetCouponByCategory(context.Context, *CouponByCategoryRequest) (*Coupons, error)
	GetGridCategory(context.Context, *emptypb.Empty) (*GridCategories, error)
	GetMyAvailableCoupon(context.Context, *emptypb.Empty) (*Coupons, error)
	GetMyCouponByStatus(context.Context, *MyCouponByStatusRequest) (*Coupons, error)
	GetSaleExplain(context.Context, *emptypb.Empty) (*SaleExplains, error)
	GetSpuByCategory(context.Context, *SpuByCategoryRequest) (*SpuPage, error)
	GetSpuById(context.Context, *SpuByIdRequest) (*SpuDetail, error)
	GetSpuLatest(context.Context, *emptypb.Empty) (*SpuPage, error)
	GetTagByType(context.Context, *TagByTypeRequest) (*Tags, error)
	GetThemeByNames(context.Context, *ThemeByNamesRequest) (*Themes, error)
	GetThemeWithSpu(context.Context, *ThemeWithSpuRequest) (*ThemeSpu, error)
	GetWholeCoupon(context.Context, *emptypb.Empty) (*Coupons, error)
	Search(context.Context, *SearchRequest) (*SpuPage, error)
}

func RegisterAppHTTPServer(s *http.Server, srv AppHTTPServer) {
	r := s.Route("/")
	r.GET("/banner/id/{id}", _App_GetBannerById0_HTTP_Handler(srv))
	r.GET("/banner/name/{name}", _App_GetBannerByName0_HTTP_Handler(srv))
	r.GET("/theme/by/names", _App_GetThemeByNames0_HTTP_Handler(srv))
	r.GET("/theme/name/{name}/with_spu", _App_GetThemeWithSpu0_HTTP_Handler(srv))
	r.GET("/activity/name/{name}", _App_GetActivityByName0_HTTP_Handler(srv))
	r.GET("/activity/name/{name}/with_coupon", _App_GetActivityWithCoupon0_HTTP_Handler(srv))
	r.GET("/coupon/by/category/{id}", _App_GetCouponByCategory0_HTTP_Handler(srv))
	r.GET("/coupon/by/whole_store", _App_GetWholeCoupon0_HTTP_Handler(srv))
	r.GET("/coupon/myself/by/status/{status}", _App_GetMyCouponByStatus0_HTTP_Handler(srv))
	r.GET("/coupon/myself/available/with_category", _App_GetMyAvailableCoupon0_HTTP_Handler(srv))
	r.POST("/coupon/collect/{id}", _App_CollectCoupon0_HTTP_Handler(srv))
	r.GET("/category/all", _App_GetAllCategory0_HTTP_Handler(srv))
	r.GET("/category/grid/all", _App_GetGridCategory0_HTTP_Handler(srv))
	r.GET("/sale_explain/fixed", _App_GetSaleExplain0_HTTP_Handler(srv))
	r.GET("/search", _App_Search0_HTTP_Handler(srv))
	r.GET("/tag/type/{type}", _App_GetTagByType0_HTTP_Handler(srv))
	r.GET("/spu/id/{id}/detail", _App_GetSpuById0_HTTP_Handler(srv))
	r.GET("/spu/latest", _App_GetSpuLatest0_HTTP_Handler(srv))
	r.GET("/spu/by/category/{id}", _App_GetSpuByCategory0_HTTP_Handler(srv))
}

func _App_GetBannerById0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BannerByIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetBannerById")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBannerById(ctx, req.(*BannerByIdRequest))
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
		var in BannerByNameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetBannerByName")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBannerByName(ctx, req.(*BannerByNameRequest))
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
		http.SetOperation(ctx, "/mall.App/GetThemeByNames")
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

func _App_GetThemeWithSpu0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ThemeWithSpuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetThemeWithSpu")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetThemeWithSpu(ctx, req.(*ThemeWithSpuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ThemeSpu)
		return ctx.Result(200, reply)
	}
}

func _App_GetActivityByName0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ActivityByNameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetActivityByName")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetActivityByName(ctx, req.(*ActivityByNameRequest))
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
		var in ActivityWithCouponRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetActivityWithCoupon")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetActivityWithCoupon(ctx, req.(*ActivityWithCouponRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ActivityCoupon)
		return ctx.Result(200, reply)
	}
}

func _App_GetCouponByCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CouponByCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetCouponByCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCouponByCategory(ctx, req.(*CouponByCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Coupons)
		return ctx.Result(200, reply)
	}
}

func _App_GetWholeCoupon0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetWholeCoupon")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWholeCoupon(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Coupons)
		return ctx.Result(200, reply)
	}
}

func _App_GetMyCouponByStatus0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MyCouponByStatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetMyCouponByStatus")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyCouponByStatus(ctx, req.(*MyCouponByStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Coupons)
		return ctx.Result(200, reply)
	}
}

func _App_GetMyAvailableCoupon0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetMyAvailableCoupon")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyAvailableCoupon(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Coupons)
		return ctx.Result(200, reply)
	}
}

func _App_CollectCoupon0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CollectCouponRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/CollectCoupon")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CollectCoupon(ctx, req.(*CollectCouponRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _App_GetAllCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetAllCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAllCategory(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AllCategory)
		return ctx.Result(200, reply)
	}
}

func _App_GetGridCategory0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetGridCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGridCategory(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GridCategories)
		return ctx.Result(200, reply)
	}
}

func _App_GetSaleExplain0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetSaleExplain")
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

func _App_Search0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SearchRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/Search")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Search(ctx, req.(*SearchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SpuPage)
		return ctx.Result(200, reply)
	}
}

func _App_GetTagByType0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TagByTypeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetTagByType")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTagByType(ctx, req.(*TagByTypeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Tags)
		return ctx.Result(200, reply)
	}
}

func _App_GetSpuById0_HTTP_Handler(srv AppHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SpuByIdRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetSpuById")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSpuById(ctx, req.(*SpuByIdRequest))
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
		http.SetOperation(ctx, "/mall.App/GetSpuLatest")
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
		var in SpuByCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.App/GetSpuByCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSpuByCategory(ctx, req.(*SpuByCategoryRequest))
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
	CollectCoupon(ctx context.Context, req *CollectCouponRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetActivityByName(ctx context.Context, req *ActivityByNameRequest, opts ...http.CallOption) (rsp *Activity, err error)
	GetActivityWithCoupon(ctx context.Context, req *ActivityWithCouponRequest, opts ...http.CallOption) (rsp *ActivityCoupon, err error)
	GetAllCategory(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *AllCategory, err error)
	GetBannerById(ctx context.Context, req *BannerByIdRequest, opts ...http.CallOption) (rsp *Banner, err error)
	GetBannerByName(ctx context.Context, req *BannerByNameRequest, opts ...http.CallOption) (rsp *Banner, err error)
	GetCouponByCategory(ctx context.Context, req *CouponByCategoryRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	GetGridCategory(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GridCategories, err error)
	GetMyAvailableCoupon(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Coupons, err error)
	GetMyCouponByStatus(ctx context.Context, req *MyCouponByStatusRequest, opts ...http.CallOption) (rsp *Coupons, err error)
	GetSaleExplain(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *SaleExplains, err error)
	GetSpuByCategory(ctx context.Context, req *SpuByCategoryRequest, opts ...http.CallOption) (rsp *SpuPage, err error)
	GetSpuById(ctx context.Context, req *SpuByIdRequest, opts ...http.CallOption) (rsp *SpuDetail, err error)
	GetSpuLatest(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *SpuPage, err error)
	GetTagByType(ctx context.Context, req *TagByTypeRequest, opts ...http.CallOption) (rsp *Tags, err error)
	GetThemeByNames(ctx context.Context, req *ThemeByNamesRequest, opts ...http.CallOption) (rsp *Themes, err error)
	GetThemeWithSpu(ctx context.Context, req *ThemeWithSpuRequest, opts ...http.CallOption) (rsp *ThemeSpu, err error)
	GetWholeCoupon(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *Coupons, err error)
	Search(ctx context.Context, req *SearchRequest, opts ...http.CallOption) (rsp *SpuPage, err error)
}

type AppHTTPClientImpl struct {
	cc *http.Client
}

func NewAppHTTPClient(client *http.Client) AppHTTPClient {
	return &AppHTTPClientImpl{client}
}

func (c *AppHTTPClientImpl) CollectCoupon(ctx context.Context, in *CollectCouponRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/coupon/collect/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/mall.App/CollectCoupon"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetActivityByName(ctx context.Context, in *ActivityByNameRequest, opts ...http.CallOption) (*Activity, error) {
	var out Activity
	pattern := "/activity/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetActivityByName"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetActivityWithCoupon(ctx context.Context, in *ActivityWithCouponRequest, opts ...http.CallOption) (*ActivityCoupon, error) {
	var out ActivityCoupon
	pattern := "/activity/name/{name}/with_coupon"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetActivityWithCoupon"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetAllCategory(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*AllCategory, error) {
	var out AllCategory
	pattern := "/category/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetAllCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetBannerById(ctx context.Context, in *BannerByIdRequest, opts ...http.CallOption) (*Banner, error) {
	var out Banner
	pattern := "/banner/id/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetBannerById"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetBannerByName(ctx context.Context, in *BannerByNameRequest, opts ...http.CallOption) (*Banner, error) {
	var out Banner
	pattern := "/banner/name/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetBannerByName"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetCouponByCategory(ctx context.Context, in *CouponByCategoryRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/by/category/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetCouponByCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetGridCategory(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GridCategories, error) {
	var out GridCategories
	pattern := "/category/grid/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetGridCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetMyAvailableCoupon(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/myself/available/with_category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetMyAvailableCoupon"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetMyCouponByStatus(ctx context.Context, in *MyCouponByStatusRequest, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/myself/by/status/{status}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetMyCouponByStatus"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetSaleExplain(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*SaleExplains, error) {
	var out SaleExplains
	pattern := "/sale_explain/fixed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetSaleExplain"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetSpuByCategory(ctx context.Context, in *SpuByCategoryRequest, opts ...http.CallOption) (*SpuPage, error) {
	var out SpuPage
	pattern := "/spu/by/category/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetSpuByCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetSpuById(ctx context.Context, in *SpuByIdRequest, opts ...http.CallOption) (*SpuDetail, error) {
	var out SpuDetail
	pattern := "/spu/id/{id}/detail"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetSpuById"))
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
	opts = append(opts, http.Operation("/mall.App/GetSpuLatest"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetTagByType(ctx context.Context, in *TagByTypeRequest, opts ...http.CallOption) (*Tags, error) {
	var out Tags
	pattern := "/tag/type/{type}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetTagByType"))
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
	opts = append(opts, http.Operation("/mall.App/GetThemeByNames"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetThemeWithSpu(ctx context.Context, in *ThemeWithSpuRequest, opts ...http.CallOption) (*ThemeSpu, error) {
	var out ThemeSpu
	pattern := "/theme/name/{name}/with_spu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetThemeWithSpu"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) GetWholeCoupon(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*Coupons, error) {
	var out Coupons
	pattern := "/coupon/by/whole_store"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/GetWholeCoupon"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AppHTTPClientImpl) Search(ctx context.Context, in *SearchRequest, opts ...http.CallOption) (*SpuPage, error) {
	var out SpuPage
	pattern := "/search"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.App/Search"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
