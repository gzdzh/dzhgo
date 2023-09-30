package service

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dict/model"
)

type DictTypeService struct {
	*dzhCore.Service
}

func NewDictTypeService() *DictTypeService {
	return &DictTypeService{
		Service: &dzhCore.Service{
			Model: model.NewDictType(),
			ListQueryOp: &dzhCore.QueryOp{
				KeyWordField: []string{"name"},
			},
		},
	}
}
