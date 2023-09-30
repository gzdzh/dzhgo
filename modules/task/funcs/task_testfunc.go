package funcs

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
)

type TaskTest struct {
}

func (t *TaskTest) Func(ctx g.Ctx, param string) error {
	g.Log().Info(ctx, "TaskTest Run ~~~~~~~~~~~~~~~~", param)
	return nil
}
func (t *TaskTest) IsSingleton() bool {
	return false
}
func (t *TaskTest) IsAllWorker() bool {
	return true
}

func init() {
	dzhCore.RegisterFunc("TaskTest", &TaskTest{})
}
