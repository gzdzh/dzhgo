package admin

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type CmsFormController struct {
	*dzhCore.Controller
}

func init() {
	var cms_form_controller = &CmsFormController{
		&dzhCore.Controller{
			Perfix:  "/admin/dzhCms/form",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsFormService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_form_controller)
}

// 增加 Welcome 演示 方法
type CmsFormWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type CmsFormWelcomeRes struct {
	*dzhCore.BaseRes
	Data interface{} `json:"data"`
}

func (c *CmsFormController) Welcome(ctx context.Context, req *CmsFormWelcomeReq) (res *CmsFormWelcomeRes, err error) {
	res = &CmsFormWelcomeRes{
		BaseRes: dzhCore.Ok("Welcome to core Admin Go"),
		Data:    gjson.New(`{"name": "core Admin Go", "age":0}`),
	}
	return
}
