package admin

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/base/api/v1"
	"github.com/gzdzh/dzhgo/modules/base/service"
)

type BaseOpen struct {
	*dzhCore.ControllerSimple
	baseSysLoginService *service.BaseSysLoginService
	baseOpenService     *service.BaseOpenService
}

func init() {
	var open = &BaseOpen{
		ControllerSimple:    &dzhCore.ControllerSimple{Perfix: "/admin/base/open"},
		baseSysLoginService: service.NewBaseSysLoginService(),
		baseOpenService:     service.NewBaseOpenService(),
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(open)
}

// 验证码接口
func (c *BaseOpen) BaseOpenCaptcha(ctx context.Context, req *v1.BaseOpenCaptchaReq) (res *dzhCore.BaseRes, err error) {
	data, err := c.baseSysLoginService.Captcha(req)
	res = dzhCore.Ok(data)
	return
}

// eps 接口请求
type BaseOpenEpsReq struct {
	g.Meta `path:"/eps" method:"GET"`
}

// eps 接口
func (c *BaseOpen) Eps(ctx context.Context, req *BaseOpenEpsReq) (res *dzhCore.BaseRes, err error) {
	if !dzhCore.Config.Eps {
		g.Log().Error(ctx, "eps is not open")
		res = dzhCore.Ok(nil)
		return
	}
	data, err := c.baseOpenService.AdminEPS(ctx)
	if err != nil {
		g.Log().Error(ctx, "eps error", err)
		return dzhCore.Fail(err.Error()), err
	}
	res = dzhCore.Ok(data)
	return
}

// login 接口
func (c *BaseOpen) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (res *dzhCore.BaseRes, err error) {
	data, err := c.baseSysLoginService.Login(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}

// RefreshTokenReq 刷新token请求
type RefreshTokenReq struct {
	g.Meta       `path:"/refreshToken" method:"GET"`
	RefreshToken string `json:"refreshToken" v:"required#refreshToken不能为空"`
}

// RefreshToken 刷新token
func (c *BaseOpen) RefreshToken(ctx context.Context, req *RefreshTokenReq) (res *dzhCore.BaseRes, err error) {
	data, err := c.baseSysLoginService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}
