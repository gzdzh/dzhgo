package app

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhCms/service"
)

type CmsSettingController struct {
	*dzhCore.Controller
}

func init() {
	var cms_setting_controller = &CmsSettingController{
		&dzhCore.Controller{
			Perfix:  "/app/dzhCms/setting",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewCmsSettingService(),
		},
	}
	// 注册路由
	dzhCore.RegisterController(cms_setting_controller)
}

// 客户列表
func (s *CmsSettingController) ContactList(ctx context.Context, req *v1.ContactListReq) (res *dzhCore.BaseRes, err error) {

	data, err := service.NewCmsSettingService().ContactList(ctx, req)
	if err != nil {
		return
	}
	res = dzhCore.Ok(data)
	return
}
