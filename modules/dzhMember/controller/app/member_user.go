package app

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhMember/service"
)

type MemberUserController struct {
	*dzhCore.Controller
}

func init() {
	var member_user_controller = &MemberUserController{
		&dzhCore.Controller{
			Perfix:  "/app/dzhMember/user",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewMemberUserService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(member_user_controller)
}
