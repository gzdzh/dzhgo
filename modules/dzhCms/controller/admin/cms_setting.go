package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsSettingController struct {
	*dzhCore.Controller
}

func init() {
	var cms_setting_controller = &CmsSettingController{
		&dzhCore.Controller{
			Perfix:  "/admin/dzhCms/setting",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsSettingService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_setting_controller)
}
