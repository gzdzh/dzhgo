package app

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsNewsController struct {
	*dzhCore.Controller
}

func init() {
	var cms_news_controller = &CmsNewsController{
		&dzhCore.Controller{
			Perfix:  "/app/dzhCms/news",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsNewsService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_news_controller)
}
