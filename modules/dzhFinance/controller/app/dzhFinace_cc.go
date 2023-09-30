package app

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
)

type DzhFinaceCcController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var dzhfinace_cc_controller = &DzhFinaceCcController{
		&dzhCore.ControllerSimple{
			Perfix: "/app/dzhFinance/cc",
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(dzhfinace_cc_controller)
}

// 增加 Welcome 演示 方法
type DzhFinaceCcWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type DzhFinaceCcWelcomeRes struct {
	*dzhCore.BaseRes
	Data interface{} `json:"data"`
}

func (c *DzhFinaceCcController) Welcome(ctx context.Context, req *DzhFinaceCcWelcomeReq) (res *DzhFinaceCcWelcomeRes, err error) {
	res = &DzhFinaceCcWelcomeRes{
		BaseRes: dzhCore.Ok("Welcome to core Admin Go"),
		Data:    gjson.New(`{"name": "core Admin Go", "age":0}`),
	}
	return
}
