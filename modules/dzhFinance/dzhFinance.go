package dzhFinance

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	_ "github.com/gzdzh/dzhgo/modules/dzhFinance/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gzdzh/dzhgo/modules/dzhFinance/controller"
	_ "github.com/gzdzh/dzhgo/modules/dzhFinance/middleware"

	baseModel "github.com/gzdzh/dzhgo/modules/base/model"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module dzhFinance init start ...")
	dzhCore.FillInitData(ctx, "dzhFinance", &baseModel.BaseSysMenu{})
	g.Log().Debug(ctx, "module dzhFinance init finished ...")
}
