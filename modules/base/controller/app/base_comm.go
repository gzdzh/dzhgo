package admin

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/service"

	"github.com/gogf/gf/v2/frame/g"
)

type BaseCommController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var base_comm_controller = &BaseCommController{
		&dzhCore.ControllerSimple{
			Perfix: "/app/base/comm",
			//    Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			//    Service: service.NewBaseCommService(),
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(base_comm_controller)
}

// eps 接口请求
type BaseCommControllerEpsReq struct {
	g.Meta `path:"/eps" method:"GET"`
}

// eps 接口
func (c *BaseCommController) Eps(ctx context.Context, req *BaseCommControllerEpsReq) (res *dzhCore.BaseRes, err error) {
	if !dzhCore.Config.Eps {
		g.Log().Error(ctx, "eps is not open")
		res = dzhCore.Ok(nil)
		return
	}
	baseOpenService := service.NewBaseOpenService()
	data, err := baseOpenService.AppEPS(ctx)
	if err != nil {
		g.Log().Error(ctx, "eps error", err)
		return dzhCore.Fail(err.Error()), err
	}
	res = dzhCore.Ok(data)
	return
}
