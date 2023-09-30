package app

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhMember/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhMember/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type MemberOpenController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var member_open_controller = &MemberOpenController{
		&dzhCore.ControllerSimple{
			Perfix: "/app/dzhMember/open",
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(member_open_controller)
}

// 增加 Welcome 演示 方法
type MemberOpenWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type MemberOpenWelcomeRes struct {
	*dzhCore.BaseRes
	Data interface{} `json:"data"`
}

func (c *MemberOpenController) Welcome(ctx context.Context, req *MemberOpenWelcomeReq) (res *MemberOpenWelcomeRes, err error) {
	res = &MemberOpenWelcomeRes{
		BaseRes: dzhCore.Ok("Welcome to core Admin Go"),
		Data:    gjson.New(`{"name": "core Admin Go", "age":0}`),
	}
	return
}

// 小程序登录
func (c *MemberOpenController) Login(ctx g.Ctx, req *v1.LoginReq) (res *dzhCore.BaseRes, err error) {

	data, err := service.NewMemberUserService().Login(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}

// 微信公众号登录
func (c *MemberOpenController) MpLoginReq(ctx g.Ctx, req *v1.MpLoginReq) (res *dzhCore.BaseRes, err error) {

	data, err := service.NewMemberUserService().MpLoginReq(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}

// 小程序登录
func (c *MemberOpenController) WxLogin(ctx g.Ctx, req *v1.WxLoginReq) (res *dzhCore.BaseRes, err error) {

	data, err := service.NewMemberUserService().WxLogin(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}

// 游客登录
func (c *MemberOpenController) TouristLogin(ctx g.Ctx, req *v1.TouristLoginReq) (res *dzhCore.BaseRes, err error) {

	data, err := service.NewMemberUserService().TouristLogin(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}

// 验证游客次数
func (c *MemberOpenController) VerifyCount(ctx g.Ctx, req *v1.VerifyCountReq) (res *dzhCore.BaseRes, err error) {

	data, err := service.NewMemberUserService().VerifyCount(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}

// 账号登录
func (c *MemberOpenController) AccountLogin(ctx g.Ctx, req *v1.AccountLoginReq) (res *dzhCore.BaseRes, err error) {

	data, err := service.NewMemberUserService().AccountLogin(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}

// 账号注册
func (c *MemberOpenController) AccountRegister(ctx g.Ctx, req *v1.AccountRegisterReq) (res *dzhCore.BaseRes, err error) {

	data, err := service.NewMemberUserService().AccountRegister(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}
