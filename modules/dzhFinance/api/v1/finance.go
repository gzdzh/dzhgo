package v1

import "github.com/gogf/gf/v2/frame/g"

type OrderPayStatusReq struct {
	g.Meta  `path:"/orderPayStatus" method:"POST"`
	Sku     *string `json:"sku"`
	TradeNo *string `json:"tradeNo" v:"required"`
}

type TipDataReq struct {
	g.Meta  `path:"/tipData" method:"POST"`
	TypeIds *[]string `json:"typeIds"`
}
