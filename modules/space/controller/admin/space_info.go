package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/space/service"
)

type SpaceInfoController struct {
	*dzhCore.Controller
}

func init() {
	var space_info_controller = &SpaceInfoController{
		&dzhCore.Controller{
			Perfix:  "/admin/space/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewSpaceInfoService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(space_info_controller)
}
