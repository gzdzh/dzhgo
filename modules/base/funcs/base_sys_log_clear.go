package funcs

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/service"
)

type BaseFuncClearLog struct {
}

// Func
func (f *BaseFuncClearLog) Func(ctx g.Ctx, param string) (err error) {
	g.Log().Info(ctx, "清理日志 BaseFuncClearLog.Func", "param", param)
	baseSysLogService := service.NewBaseSysLogService()
	if param == "true" {
		err = baseSysLogService.Clear(true)
	} else {
		err = baseSysLogService.Clear(false)
	}
	return
}

// IsSingleton
func (f *BaseFuncClearLog) IsSingleton() bool {
	return true
}

// IsAllWorker
func (f *BaseFuncClearLog) IsAllWorker() bool {
	return false
}

// init
func init() {
	dzhCore.RegisterFunc("BaseFuncClearLog", &BaseFuncClearLog{})
}
