package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/space/service"
)

type SpaceTypeController struct {
	*dzhCore.Controller
}

func init() {
	var space_type_controller = &SpaceTypeController{
		&dzhCore.Controller{
			Perfix:  "/admin/space/type",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewSpaceTypeService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(space_type_controller)
}
