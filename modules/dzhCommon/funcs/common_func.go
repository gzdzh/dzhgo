package funcs

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
)

type CommonFuncSent struct {
}

// Func
// 定时任务发送通知
func (f *CommonFuncSent) Func(ctx g.Ctx, param string) (err error) {
	g.Log().Debug(ctx, "发送 CommonFuncSent", "param", param)

	// if param == "true" {
	// _, _ = orderService.ServiceSentMail(ctx, true)
	// }

	return
}

// IsSingleton
func (f *CommonFuncSent) IsSingleton() bool {
	return true
}

// IsAllWorker
func (f *CommonFuncSent) IsAllWorker() bool {
	return false
}

// init
func init() {
	dzhCore.RegisterFunc("CommonFuncSent", &CommonFuncSent{})
}
