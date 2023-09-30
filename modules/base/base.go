package base

import (
	_ "github.com/gzdzh/dzhgo/modules/base/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/gzdzh/dzhgo/modules/base/controller"
	_ "github.com/gzdzh/dzhgo/modules/base/funcs"
	_ "github.com/gzdzh/dzhgo/modules/base/middleware"
	"github.com/gzdzh/dzhgo/modules/base/model"

	"github.com/gzdzh/dzhgo/dzhCore"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module base init start ...")

	dzhCore.FillInitData(ctx, "base", &model.BaseSysMenu{})
	dzhCore.FillInitData(ctx, "base", &model.BaseSysUser{})
	dzhCore.FillInitData(ctx, "base", &model.BaseSysUserRole{})
	dzhCore.FillInitData(ctx, "base", &model.BaseSysRole{})
	dzhCore.FillInitData(ctx, "base", &model.BaseSysRoleMenu{})
	dzhCore.FillInitData(ctx, "base", &model.BaseSysDepartment{})
	dzhCore.FillInitData(ctx, "base", &model.BaseSysRoleDepartment{})
	dzhCore.FillInitData(ctx, "base", &model.BaseSysParam{})

	g.Log().Debug(ctx, "module base init finished ...")

}
