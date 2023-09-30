package admin

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dict/service"

	"github.com/gogf/gf/v2/frame/g"
)

type DictInfoController struct {
	*dzhCore.Controller
}

func init() {
	var dict_info_controller = &DictInfoController{
		&dzhCore.Controller{
			Perfix:  "/admin/dict/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDictInfoService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(dict_info_controller)
}

// Data 方法请求
type DictInfoDataReq struct {
	g.Meta `path:"/data" method:"POST"`
	Types  []string `json:"types"`
}

// Data 方法 获得字典数据
func (c *DictInfoController) Data(ctx context.Context, req *DictInfoDataReq) (res *dzhCore.BaseRes, err error) {
	service := service.NewDictInfoService()
	data, err := service.Data(ctx, req.Types)
	res = dzhCore.Ok(data)
	return
}
