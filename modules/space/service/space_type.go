package service

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/space/model"
)

type SpaceTypeService struct {
	*dzhCore.Service
}

func NewSpaceTypeService() *SpaceTypeService {
	return &SpaceTypeService{
		&dzhCore.Service{
			Model: model.NewSpaceType(),
		},

		// Service: dzhCore.NewService(model.NewSpaceType()),
	}
}
