package biz

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
)
type Coupon struct {
	Id int64 
	StartTime time.Time
	EndTime time.Time
	WholeStore int8 `json:"whole_store,omitempty"`
	Category []Category `json:"category,omitempty"`
	Type int `json:"type,omitempty"`
	FullMoney float64 `json:"full_money,omitempty"`
	// Minus holds the value of the "minus" field.
	Minus float64 `json:"minus,omitempty"`
	// Rate holds the value of the "rate" field.
	// 国内多是打折，国外多是百分比 off
	Rate float64 `json:"rate,omitempty"`
}
type Category struct {
	Id int64 
}
type CouponRepo interface {
}
type CouponUsecase struct {
	repo CouponRepo
	log  *log.Helper
}

func NewCouponUsecase(repo CouponRepo, logger log.Logger) *CouponUsecase {
	return &CouponUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
