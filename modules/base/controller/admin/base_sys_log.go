package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/service"
)

type BaseSysLogController struct {
	*dzhCore.Controller
}

func init() {
	var base_sys_log_controller = &BaseSysLogController{
		&dzhCore.Controller{
			Perfix:  "/admin/base/sys/log",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysLogService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(base_sys_log_controller)
}

// SetKeepReq
type SetKeepReq struct {
	g.Meta `method:"POST" path:"/setKeep" summary:"设置保留天数" tags:"系统日志"`
	Value  int `json:"value" v:"required#请输入保留天数"`
}

// SetKeep 设置保留天数
func (c *BaseSysLogController) SetKeep(ctx g.Ctx, req *SetKeepReq) (res *dzhCore.BaseRes, err error) {
	var (
		BaseSysConfService = service.NewBaseSysConfService()
	)
	err = BaseSysConfService.UpdateValue("logKeep", gconv.String(req.Value))
	return
}

// GetKeepReq
type GetKeepReq struct {
	g.Meta `method:"GET" path:"/getKeep" summary:"获取保留天数" tags:"系统日志"`
}

// GetKeep 获取保留天数
func (c *BaseSysLogController) GetKeep(ctx g.Ctx, req *GetKeepReq) (res *dzhCore.BaseRes, err error) {
	var (
		BaseSysConfService = service.NewBaseSysConfService()
	)
	// res.Data = BaseSysConfService.GetValue("logKeep")
	res = dzhCore.Ok(BaseSysConfService.GetValue("logKeep"))
	return
}

// ClearReq
type ClearReq struct {
	g.Meta `method:"POST" path:"/clear" summary:"清空日志" tags:"系统日志"`
}

// Clear 清空日志
func (c *BaseSysLogController) Clear(ctx g.Ctx, req *ClearReq) (res *dzhCore.BaseRes, err error) {
	var (
		BaseSysLogService = service.NewBaseSysLogService()
	)
	err = BaseSysLogService.Clear(true)
	return
}
