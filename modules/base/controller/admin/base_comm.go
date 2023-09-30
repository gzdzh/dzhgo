package admin

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/service"
)

type BaseCommController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var base_comm_controller = &BaseCommController{
		ControllerSimple: &dzhCore.ControllerSimple{
			Perfix: "/admin/base/comm",
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(base_comm_controller)
}

// BaseCommPersonReq 接口请求参数
type BaseCommPersonReq struct {
	g.Meta        `path:"/person" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// Person 方法
func (c *BaseCommController) Person(ctx context.Context, req *BaseCommPersonReq) (res *dzhCore.BaseRes, err error) {
	var (
		baseSysUserService = service.NewBaseSysUserService()
		admin              = dzhCore.GetAdmin(ctx)
	)
	data, err := baseSysUserService.Person(admin.UserId)
	res = dzhCore.Ok(data)
	return
}

// BaseCommPermmenuReq 接口请求参数
type BaseCommPermmenuReq struct {
	g.Meta        `path:"/permmenu" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// Permmenu 方法
func (c *BaseCommController) Permmenu(ctx context.Context, req *BaseCommPermmenuReq) (res *dzhCore.BaseRes, err error) {

	var (
		baseSysPermsService = service.NewBaseSysPermsService()
		admin               = dzhCore.GetAdmin(ctx)
	)
	res = dzhCore.Ok(baseSysPermsService.Permmenu(ctx, admin.RoleIds))
	return
}

type BaseCommLogoutReq struct {
	g.Meta        `path:"/logout" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

// Logout BaseCommLogout 方法
func (c *BaseCommController) Logout(ctx context.Context, req *BaseCommLogoutReq) (res *dzhCore.BaseRes, err error) {
	var (
		BaseSysLoginService = service.NewBaseSysLoginService()
	)
	err = BaseSysLoginService.Logout(ctx)
	res = dzhCore.Ok(nil)
	return
}

type BaseCommUploadModeReq struct {
	g.Meta        `path:"/uploadMode" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// UploadMode 方法
func (c *BaseCommController) UploadMode(ctx context.Context, req *BaseCommUploadModeReq) (res *dzhCore.BaseRes, err error) {
	data, err := dzhCore.File().GetMode()
	res = dzhCore.Ok(data)
	return
}

type BaseCommUploadReq struct {
	g.Meta        `path:"/upload" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

// Upload 方法
func (c *BaseCommController) Upload(ctx context.Context, req *BaseCommUploadReq) (res *dzhCore.BaseRes, err error) {
	data, err := dzhCore.File().Upload(ctx)
	res = dzhCore.Ok(data)
	return
}

type PersonUpdateReq struct {
	g.Meta        `path:"/personUpdate" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

// PersonUpdate 方法
func (c *BaseCommController) PersonUpdate(ctx g.Ctx, req *PersonUpdateReq) (res *dzhCore.BaseRes, err error) {
	var (
		baseSysUserService = service.NewBaseSysUserService()
	)

	_, err = baseSysUserService.ServiceUpdate(ctx, &dzhCore.UpdateReq{})
	if err != nil {
		return
	}

	res = dzhCore.Ok(nil)
	return
}
