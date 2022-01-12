package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewThemeUsecase,
	NewBannerUsecase, NewActivityUsecase, NewCouponUsecase, NewSkuUsecase, NewSpuUsecase, NewOrderUsecase,
)
