package admin

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCommon/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type CommonSettingController struct {
	*dzhCore.Controller
}

func init() {
	var common_setting_controller = &CommonSettingController{
		&dzhCore.Controller{
			Perfix:  "/admin/common/setting",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCommonSettingService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(common_setting_controller)
}

// 增加 Welcome 演示 方法
type CommonSettingWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type CommonSettingWelcomeRes struct {
	*dzhCore.BaseRes
	Data interface{} `json:"data"`
}

func (c *CommonSettingController) Welcome(ctx context.Context, req *CommonSettingWelcomeReq) (res *CommonSettingWelcomeRes, err error) {
	res = &CommonSettingWelcomeRes{
		BaseRes: dzhCore.Ok("Welcome to core Admin Go"),
		Data:    gjson.New(`{"name": "core Admin Go", "age":0}`),
	}
	return
}
