package service

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"
)

type CmsGoodTypeService struct {
	*dzhCore.Service
}

func NewCmsGoodTypeService() *CmsGoodTypeService {
	return &CmsGoodTypeService{
		&dzhCore.Service{
			Model: model.NewCmsGoodType(),
		},
	}
}
