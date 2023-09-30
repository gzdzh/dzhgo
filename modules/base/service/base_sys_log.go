package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/model"
)

type BaseSysLogService struct {
	*dzhCore.Service
}

func NewBaseSysLogService() *BaseSysLogService {
	return &BaseSysLogService{
		&dzhCore.Service{
			Model: model.NewBaseSysLog(),
			PageQueryOp: &dzhCore.QueryOp{
				KeyWordField: []string{"name", "params", "ipAddr"},
				Select:       "base_sys_log.*,user.name ",
				Join: []*dzhCore.JoinOp{
					{
						Model:     model.NewBaseSysUser(),
						Alias:     "user",
						Type:      "LeftJoin",
						Condition: "user.id = base_sys_log.userID",
					},
				},
			},
		},
	}
}

// Record 记录日志
func (s *BaseSysLogService) Record(ctx g.Ctx) {
	var (
		admin = dzhCore.GetAdmin(ctx)
		r     = g.RequestFromCtx(ctx)
	)
	baseSysLog := model.NewBaseSysLog()
	baseSysLog.UserID = admin.UserId
	baseSysLog.Action = r.Method + ":" + r.URL.Path
	baseSysLog.IP = r.GetClientIp()
	baseSysLog.IPAddr = r.GetClientIp()
	baseSysLog.Params = r.GetBodyString()
	m := dzhCore.DBM(s.Model)
	m.Insert(g.Map{
		"userId": baseSysLog.UserID,
		"action": baseSysLog.Action,
		"ip":     baseSysLog.IP,
		"ipAddr": baseSysLog.IPAddr,
		"params": baseSysLog.Params,
	})
}

// Clear 清除日志
func (s *BaseSysLogService) Clear(isAll bool) (err error) {
	BaseSysConfService := NewBaseSysConfService()
	m := dzhCore.DBM(s.Model)
	if isAll {
		_, err = m.Delete("1=1")
	} else {
		keepDays := gconv.Int(BaseSysConfService.GetValue("logKeep"))
		_, err = m.Delete("createTime < ?", gtime.Now().AddDate(0, 0, -keepDays).String())
	}
	return
}
