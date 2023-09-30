package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/service"
)

type BaseSysMenuController struct {
	*dzhCore.Controller
}

func init() {
	var base_sys_menu_controller = &BaseSysMenuController{
		&dzhCore.Controller{
			Perfix:  "/admin/base/sys/menu",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysMenuService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(base_sys_menu_controller)
}
