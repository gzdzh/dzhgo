package service

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCommon/model"
)

type CommonSettingService struct {
	*dzhCore.Service
}

func NewCommonSettingService() *CommonSettingService {
	return &CommonSettingService{
		&dzhCore.Service{
			Model: model.NewCommonSetting(),
		},
	}
}
