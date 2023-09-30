package dzhMember

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	_ "github.com/gzdzh/dzhgo/modules/dzhMember/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gzdzh/dzhgo/modules/dzhMember/controller"
	_ "github.com/gzdzh/dzhgo/modules/dzhMember/middleware"

	baseModel "github.com/gzdzh/dzhgo/modules/base/model"
	"github.com/gzdzh/dzhgo/modules/dzhMember/model"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module dzhMember init start ...")
	dzhCore.FillInitData(ctx, "dzhMember", &model.MemberUser{})
	dzhCore.FillInitData(ctx, "dzhMember", &baseModel.BaseSysMenu{})
	g.Log().Debug(ctx, "module dzhMember init finished ...")
}
