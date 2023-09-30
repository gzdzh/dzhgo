package dict

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gzdzh/dzhgo/dzhCore"
	_ "github.com/gzdzh/dzhgo/modules/dict/packed"

	_ "github.com/gzdzh/dzhgo/modules/dict/controller"
	"github.com/gzdzh/dzhgo/modules/dict/model"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module dict init start ...")
	dzhCore.FillInitData(ctx, "dict", &model.DictInfo{})
	dzhCore.FillInitData(ctx, "dict", &model.DictType{})
	g.Log().Debug(ctx, "module dict init finished ...")
}
