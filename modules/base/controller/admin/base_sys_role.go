package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/service"
)

type BaseSysRoleController struct {
	*dzhCore.Controller
}

func init() {
	var base_sys_role_controller = &BaseSysRoleController{
		&dzhCore.Controller{
			Perfix:  "/admin/base/sys/role",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysRoleService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(base_sys_role_controller)
}
