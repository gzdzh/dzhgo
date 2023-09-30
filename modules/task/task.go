package demo

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	_ "github.com/gzdzh/dzhgo/modules/task/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/gzdzh/dzhgo/modules/task/controller"
	_ "github.com/gzdzh/dzhgo/modules/task/funcs"
	_ "github.com/gzdzh/dzhgo/modules/task/middleware"
	"github.com/gzdzh/dzhgo/modules/task/model"
)

func init() {
	var (
		taskInfo = model.NewTaskInfo()
		ctx      = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module task init start ...")
	dzhCore.FillInitData(ctx, "task", taskInfo)

	result, err := dzhCore.DBM(taskInfo).Where("status = ?", 1).All()
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		id := v["id"].String()
		dzhCore.RunFunc(ctx, "TaskAddTask("+id+")")
	}
	g.Log().Debug(ctx, "module task init finished ...")

}
