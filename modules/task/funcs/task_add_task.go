package funcs

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/task/model"
	"github.com/gzdzh/dzhgo/modules/task/service"
)

type TaskAddTask struct {
}

func (t *TaskAddTask) Func(ctx g.Ctx, id string) error {
	taskInfo := model.NewTaskInfo()
	result, err := dzhCore.DBM(taskInfo).Where("id = ?", id).One()
	if err != nil {
		return err
	}
	if result["taskType"].Int() == 1 {
		every := result["every"].Uint() / 1000
		cron := "@every " + gconv.String(every) + "s"
		funcstring := result["service"].String()
		startDate := result["startDate"].String()
		err = service.EnableTask(ctx, id, funcstring, cron, startDate)
	} else {
		cron := result["cron"].String()
		funcstring := result["service"].String()
		startDate := result["startDate"].String()
		err = service.EnableTask(ctx, id, funcstring, cron, startDate)
	}

	return err
}

func (t *TaskAddTask) IsSingleton() bool {
	return false
}

func (t *TaskAddTask) IsAllWorker() bool {
	return true
}

func init() {
	dzhCore.RegisterFunc("TaskAddTask", &TaskAddTask{})
}
