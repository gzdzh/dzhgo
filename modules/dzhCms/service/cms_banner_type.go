package service

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"
)

type CmsBannerTypeService struct {
	*dzhCore.Service
}

func NewCmsBannerTypeService() *CmsBannerTypeService {
	return &CmsBannerTypeService{
		&dzhCore.Service{
			Model: model.NewCmsBannerType(),
		},
	}
}
