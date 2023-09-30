package model

type PayModel struct {
	SkuType string  `json:"skuType" v:"required"`
	TradeNo string  `json:"tradeNo" v:"required"`
	Price   float64 `json:"price" v:"required"`
	Name    string  `json:"name" v:"required"`
}
