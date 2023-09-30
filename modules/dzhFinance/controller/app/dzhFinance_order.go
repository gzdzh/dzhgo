package app

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhFinance/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhFinance/service"

	"github.com/gogf/gf/v2/frame/g"
)

type DzhFinanceOrderController struct {
	*dzhCore.Controller
}

func init() {
	var dzhfinance_order_controller = &DzhFinanceOrderController{
		&dzhCore.Controller{
			Perfix:  "/app/dzhFinance/order",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDzhFinanceOrderService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(dzhfinance_order_controller)
}

// 查询订单支付状态
func (s *DzhFinanceOrderController) OrderPayStatus(ctx g.Ctx, req *v1.OrderPayStatusReq) (res *dzhCore.BaseRes, err error) {
	data, err := service.NewDzhFinanceOrderService().OrderPayStatus(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}
