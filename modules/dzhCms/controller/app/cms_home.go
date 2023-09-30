package app

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsHomeController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var cms_home_controller = &CmsHomeController{
		&dzhCore.ControllerSimple{
			Perfix: "/app/dzhCms/home",
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(cms_home_controller)
}

func (c *CmsHomeController) Home(ctx context.Context, req *v1.HomeReq) (res *dzhCore.BaseRes, err error) {
	data, err := service.NewCmsHomeService().Home(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}
