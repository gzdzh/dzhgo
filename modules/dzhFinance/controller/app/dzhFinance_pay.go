package app

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhFinance/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhFinance/service"
)

type DzhFinancePayController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var dzhfinance_pay_controller = &DzhFinancePayController{
		&dzhCore.ControllerSimple{
			Perfix: "/app/dzhfinance/pay",
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(dzhfinance_pay_controller)
}

// 服务商native下单
func (c *DzhFinancePayController) V3Native(ctx context.Context, req *v1.V3PartnerNativeReq) (res *dzhCore.BaseRes, err error) {
	data, err := service.NewDzhFinancePayService().V3Native(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}

// native回调
func (c *DzhFinancePayController) PayNotice(ctx context.Context, req *v1.NoticeReq) (res *dzhCore.BaseRes, err error) {
	data, err := service.NewDzhFinancePayService().PayNotice(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}
