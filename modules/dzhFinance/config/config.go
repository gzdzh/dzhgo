package config

import (
	"github.com/gogf/gf/v2/frame/g"
)

func WxPay(ctx g.Ctx) g.Map {
	wxPayMap := g.Cfg().MustGetWithEnv(ctx, "wxPay").Map()
	return wxPayMap
}
