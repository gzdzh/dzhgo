package service

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type CmsSettingService struct {
	*dzhCore.Service
}

// 客户列表
func (s *CmsSettingService) ContactList(ctx context.Context, req *v1.ContactListReq) (data interface{}, err error) {

	list, err := dzhCore.DBM(s.Model).Where("id", 1).Value("contactList")

	j, _ := gjson.LoadJson(list)

	data = j
	return
}

func NewCmsSettingService() *CmsSettingService {
	return &CmsSettingService{
		&dzhCore.Service{
			Model: model.NewCmsSetting(),
		},
	}
}
