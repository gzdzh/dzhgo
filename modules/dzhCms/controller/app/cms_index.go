package app

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gview"
)

type CmsIndexController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var cms_index_controller = &CmsIndexController{
		&dzhCore.ControllerSimple{
			Perfix: "/app/dzhCms/index",
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(cms_index_controller)
}

func (c *CmsIndexController) Index(ctx context.Context, req *v1.IndexReq) (res *dzhCore.BaseRes, err error) {

	view := gview.New()
	// 设置模板目录
	view.SetPath("template")
	view.Assigns(map[string]interface{}{
		"name":  "john",
		"age":   18,
		"score": 100,
	})
	// 渲染模板
	content, err := view.Parse(ctx, "index.tpl")
	if err != nil {
		panic(err)
	}

	var r = g.RequestFromCtx(ctx)
	r.Response.Writefln(content)

	return
}

func (c *CmsIndexController) Index2(ctx context.Context, req *v1.Index2Req) (res *dzhCore.BaseRes, err error) {

	var r = g.RequestFromCtx(ctx)
	r.Response.WriteTpl("index.tpl", g.Map{
		"name":  "姓名",
		"age":   18,
		"score": 100,
	})
	return
}
