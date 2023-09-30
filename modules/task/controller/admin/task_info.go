package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/task/service"
)

type TaskInfoController struct {
	*dzhCore.Controller
}

func init() {
	var task_info_controller = &TaskInfoController{
		&dzhCore.Controller{
			Perfix:  "/admin/task/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page", "Start", "Stop"},
			Service: service.NewTaskInfoService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(task_info_controller)
}

// TaskInfoStopReq 请求参数
type TaskInfoStopReq struct {
	g.Meta `path:"/stop" method:"GET"`
	ID     int64 `json:"id" v:"required#请输入id"`
}

// Stop 停止任务
func (c *TaskInfoController) Stop(ctx g.Ctx, req *TaskInfoStopReq) (res *dzhCore.BaseRes, err error) {

	err = dzhCore.ClusterRunFunc(ctx, "TaskStopFunc("+gconv.String(req.ID)+")")
	if err != nil {
		return dzhCore.Fail(err.Error()), err
	}
	res = dzhCore.Ok("停止成功")
	return
}

// TaskInfoStartReq 请求参数
type TaskInfoStartReq struct {
	g.Meta `path:"/start" method:"GET"`
	ID     int64 `json:"id" v:"required#请输入id"`
}

// Start 启动任务
func (c *TaskInfoController) Start(ctx g.Ctx, req *TaskInfoStartReq) (res *dzhCore.BaseRes, err error) {

	err = dzhCore.ClusterRunFunc(ctx, "TaskStartFunc("+gconv.String(req.ID)+")")
	if err != nil {
		return dzhCore.Fail(err.Error()), err
	}
	res = dzhCore.Ok("启动成功")
	return
}

// TaskInfoOnceReq 请求参数
type TaskInfoOnceReq struct {
	g.Meta `path:"/once" method:"POST"`
	ID     int64 `json:"id" v:"required#请输入id"`
}

// Once 执行一次
func (c *TaskInfoController) Once(ctx g.Ctx, req *TaskInfoOnceReq) (res *dzhCore.BaseRes, err error) {
	err = c.Service.(*service.TaskInfoService).Once(ctx, req.ID)
	if err != nil {
		return dzhCore.Fail(err.Error()), err
	}
	res = dzhCore.Ok("执行成功")
	return
}

// TaskInfoLogReq 请求参数
type TaskInfoLogReq struct {
	g.Meta `path:"/log" method:"GET"`
	ID     int64 `json:"id"`
	Status int   `json:"status"`
}

// Log 任务日志
func (c *TaskInfoController) Log(ctx g.Ctx, req *TaskInfoLogReq) (res *dzhCore.BaseRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	param := r.GetQueryMapStrStr()
	data, err := c.Service.(*service.TaskInfoService).Log(ctx, param)
	if err != nil {
		return dzhCore.Fail(err.Error()), err
	}
	res = dzhCore.Ok(data)
	return
}
