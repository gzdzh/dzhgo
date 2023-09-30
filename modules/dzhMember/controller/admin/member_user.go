package admin

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhMember/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberUserController struct {
	*dzhCore.Controller
}

func init() {
	var member_user_controller = &MemberUserController{
		&dzhCore.Controller{
			Perfix:  "/admin/dzhMember/user",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewMemberUserService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(member_user_controller)
}

// 增加 Welcome 演示 方法
type MemberUserWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type MemberUserWelcomeRes struct {
	*dzhCore.BaseRes
	Data interface{} `json:"data"`
}

func (c *MemberUserController) Welcome(ctx context.Context, req *MemberUserWelcomeReq) (res *MemberUserWelcomeRes, err error) {
	res = &MemberUserWelcomeRes{
		BaseRes: dzhCore.Ok("Welcome to core Admin Go"),
		Data:    gjson.New(`{"name": "core Admin Go", "age":0}`),
	}
	return
}
