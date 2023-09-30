package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsGoodNameController struct {
	*dzhCore.Controller
}

func init() {
	var cms_good_name_controller = &CmsGoodNameController{
		&dzhCore.Controller{
			Perfix:  "/admin/dzhCms/goodname",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsGoodNameService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_good_name_controller)
}
