package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsBannerTypeController struct {
	*dzhCore.Controller
}

func init() {
	var cms_banner_type_controller = &CmsBannerTypeController{
		&dzhCore.Controller{
			Perfix:  "/admin/dzhCms/bannerType",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsBannerTypeService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_banner_type_controller)
}
