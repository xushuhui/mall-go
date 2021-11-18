package logic

import (
	"github.com/xushuhui/goal/core"
	"mall-go/internal/data"

	"mall-go/internal/request"
)

type CouponChecker struct {
}
type OrderChecker struct {
	orderDTO      request.PlaceOrder
	serverSkuList []data.Sku
	couponChecker CouponChecker
	maxSkuLimit   int
	orderSkuList  []data.OrderSku
}

func NewOrderChecker(req request.PlaceOrder, serverSkuList []data.Sku, checker CouponChecker, maxSkuLimit int) *OrderChecker {
	return &OrderChecker{
		orderDTO:      req,
		serverSkuList: serverSkuList,
		couponChecker: checker,
		maxSkuLimit:   maxSkuLimit,
	}
}
func (o *OrderChecker) GetLeaderImg() string {
	return o.serverSkuList[0].Img
}
func (o *OrderChecker) GetLeaderTitle() string {
	return o.serverSkuList[0].Title
}
func (o *OrderChecker) GetTotalCount() (totalCount int) {

	for _, v := range o.orderDTO.SkuInfoList {
		totalCount = totalCount + v.Count
	}
	return
}

func (o *OrderChecker) IsOk() (err error) {
	var serverTotalPrice float64
	err = skuNotOnSale(len(o.orderDTO.SkuInfoList), len(o.serverSkuList))
	if err != nil {
		return
	}
	for i := 0; i < len(o.serverSkuList); i++ {
		sku := o.serverSkuList[i]
		skuInfoDTO := o.orderDTO.SkuInfoList[i]
		err = containsSoldOutSku(sku)
		if err != nil {
			return err
		}
		err = beyondSkuStock(sku, skuInfoDTO)
		if err != nil {
			return err
		}
		err = o.beyondMaxSkuLimit(skuInfoDTO)
		if err != nil {
			return err
		}
		price, err := calculateSkuOrderPrice(skuInfoDTO, sku)
		if err != nil {
			return err
		}
		serverTotalPrice = serverTotalPrice + price
		o.orderSkuList = append(o.orderSkuList, data.NewOrderSku(sku, skuInfoDTO))
		//skuOrderBOList.add(new SkuOrderBO(sku, skuInfoDTO));

	}
	err = totalPriceIsOk(o.orderDTO.TotalPrice, serverTotalPrice)
	if err != nil {
		return err
	}
	if o.couponChecker != struct{}{} {

	}
	return
}
func totalPriceIsOk(orderTotalPrice float64, serverTotalPrice float64) (err error) {
	if orderTotalPrice != serverTotalPrice {
		err = core.ParamsError(core.InvalidParams)
		return
	}
	return
}
func calculateSkuOrderPrice(skuInfoDTO request.SkuInfo, sku data.Sku) (price float64, err error) {
	if skuInfoDTO.Count <= 0 {
		err = core.ParamsError(core.InvalidParams)
		return
	}
	price = sku.GetActualPrice() * float64(skuInfoDTO.Count)
	return
}
func skuNotOnSale(count1, count2 int) (err error) {
	if count1 != count2 {
		err = core.ParamsError(core.InvalidParams)
	}
	return
}
func containsSoldOutSku(sku data.Sku) (err error) {
	if sku.Stock == 0 {
		err = core.ParamsError(core.InvalidParams)
	}
	return
}
func beyondSkuStock(sku data.Sku, skuInfoDTO request.SkuInfo) (err error) {
	if sku.Stock < skuInfoDTO.Count {
		err = core.ParamsError(core.InvalidParams)
	}
	return
}
func (o *OrderChecker) beyondMaxSkuLimit(skuInfoDTO request.SkuInfo) (err error) {
	if skuInfoDTO.Count > o.maxSkuLimit {
		err = core.ParamsError(core.InvalidParams)
	}
	return
}
