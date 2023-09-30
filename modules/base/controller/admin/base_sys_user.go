package admin

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/service"
)

type BaseSysUserController struct {
	*dzhCore.Controller
}

func init() {
	var base_sys_user_controller = &BaseSysUserController{
		&dzhCore.Controller{
			Perfix:  "/admin/base/sys/user",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page", "Move"},
			Service: service.NewBaseSysUserService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(base_sys_user_controller)
}

type UserMoveReq struct {
	g.Meta        `path:"/move" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

func (c *BaseSysUserController) Move(ctx context.Context, req *UserMoveReq) (res *dzhCore.BaseRes, err error) {
	err = service.NewBaseSysUserService().Move(ctx)
	res = dzhCore.Ok(nil)
	return
}
