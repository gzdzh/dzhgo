package funcs

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/task/model"
	"github.com/gzdzh/dzhgo/modules/task/service"
)

type TaskStopFunc struct {
}

func (t *TaskStopFunc) Func(ctx g.Ctx, id string) error {
	taskInfo := model.NewTaskInfo()
	_, err := dzhCore.DBM(taskInfo).Where("id = ?", id).Update(g.Map{"status": 0})
	if err != nil {
		return err
	}

	return service.DisableTask(ctx, id)
}

func (t *TaskStopFunc) IsSingleton() bool {
	return false
}

func (t *TaskStopFunc) IsAllWorker() bool {
	return true
}

func init() {
	dzhCore.RegisterFunc("TaskStopFunc", &TaskStopFunc{})
}
