package config

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"
	"github.com/gzdzh/dzhgo/modules/dzhCommon/utility/config"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

// 数据库配置文件
func GetSetting() (data gdb.Record) {
	setting, _ := dzhCore.DBM(model.NewCmsSetting()).Where("id", 1).One()
	return setting
}

var BDTRAN = func(ctx context.Context) map[string]string {

	setting := GetSetting()

	if gconv.String(setting.GMap().Get("baiduTranApiKey")) != "" && gconv.String(setting.GMap().Get("baiduTranSecretKey")) != "" {
		glog.Debugf(ctx, "baiduTranApiKey :%v, baiduTranSecretKey: %v", gconv.String(setting.GMap().Get("baiduTranApiKey")), gconv.String(setting.GMap().Get("baiduTranSecretKey")))
		return g.MapStrStr{
			"BD_API_KEY":    gconv.String(setting.GMap().Get("baiduTranApiKey")),
			"BD_SECRET_KEY": gconv.String(setting.GMap().Get("baiduTranSecretKey")),
		}
	} else {
		return config.GetMapStrStr(ctx, "baiduTran")
	}

}
