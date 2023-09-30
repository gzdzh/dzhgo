package admin

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type CommonSentController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var common_sms_controller = &CommonSentController{
		&dzhCore.ControllerSimple{
			Perfix: "/admin/demo/common_sms",
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(common_sms_controller)
}

// 增加 Welcome 演示 方法
type CommonSentWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type CommonSentWelcomeRes struct {
	*dzhCore.BaseRes
	Data interface{} `json:"data"`
}

func (c *CommonSentController) Welcome(ctx context.Context, req *CommonSentWelcomeReq) (res *CommonSentWelcomeRes, err error) {
	res = &CommonSentWelcomeRes{
		BaseRes: dzhCore.Ok("Welcome to core Admin Go"),
		Data:    gjson.New(`{"name": "core Admin Go", "age":0}`),
	}
	return
}
