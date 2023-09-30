package consts

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"

	"github.com/gogf/gf/v2/util/gconv"
)

func GetFilter() string {

	info, _ := dzhCore.DBM(model.NewCmsSetting()).Where("id", 1).One()

	return gconv.String(info["phrase"])
}
