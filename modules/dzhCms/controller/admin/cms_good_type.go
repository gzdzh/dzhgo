package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsGoodTypeController struct {
	*dzhCore.Controller
}

func init() {
	var cms_good_type_controller = &CmsGoodTypeController{
		&dzhCore.Controller{
			Perfix:  "/admin/dzhCms/goodtype",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsGoodTypeService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_good_type_controller)
}
