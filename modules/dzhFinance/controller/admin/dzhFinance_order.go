package admin

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhFinance/service"
)

type DzhFinanceOrderController struct {
	*dzhCore.Controller
}

func init() {
	var dzhfinance_order_controller = &DzhFinanceOrderController{
		&dzhCore.Controller{
			Perfix:  "/admin/dzhFinance/order",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDzhFinanceOrderService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(dzhfinance_order_controller)
}
