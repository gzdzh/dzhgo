package app

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsBannerController struct {
	*dzhCore.Controller
}

func init() {
	var cms_banner_controller = &CmsBannerController{
		&dzhCore.Controller{
			Perfix:  "/app/dzhCms/banner",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsBannerService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_banner_controller)
}

func (c *CmsBannerController) BannerList(ctx context.Context, req *v1.BannerListReq) (res *dzhCore.BaseRes, err error) {
	data, err := service.NewCmsBannerService().BannerList(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}
