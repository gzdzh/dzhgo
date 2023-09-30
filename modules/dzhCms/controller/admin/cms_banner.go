package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsBannerController struct {
	*dzhCore.Controller
}

func init() {
	var cms_banner_controller = &CmsBannerController{
		&dzhCore.Controller{
			Perfix:  "/admin/dzhCms/banner",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsBannerService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_banner_controller)
}
