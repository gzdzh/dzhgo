package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dict/service"
)

type DictTypeController struct {
	*dzhCore.Controller
}

func init() {
	var dict_type_controller = &DictTypeController{
		&dzhCore.Controller{
			Perfix:  "/admin/dict/type",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDictTypeService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(dict_type_controller)
}
