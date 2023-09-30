package dzhCms

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	_ "github.com/gzdzh/dzhgo/modules/dzhCms/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gzdzh/dzhgo/modules/dzhCms/controller"
	_ "github.com/gzdzh/dzhgo/modules/dzhCms/middleware"

	baseModel "github.com/gzdzh/dzhgo/modules/base/model"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module dzhCms init start ...")
	dzhCore.FillInitData(ctx, "dzhCms", &model.CmsSetting{})
	dzhCore.FillInitData(ctx, "dzhCms", &baseModel.BaseSysMenu{})
	g.Log().Debug(ctx, "module dzhCms init finished ...")
}
