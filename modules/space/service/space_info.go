package service

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/space/model"
)

type SpaceInfoService struct {
	*dzhCore.Service
}

func NewSpaceInfoService() *SpaceInfoService {
	return &SpaceInfoService{
		&dzhCore.Service{
			Model: model.NewSpaceInfo(),
		},

		// Service: dzhCore.NewService(model.NewSpaceInfo()),
	}
}
