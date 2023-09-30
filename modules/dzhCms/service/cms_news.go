package service

import (
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"
)

type CmsNewsService struct {
	*dzhCore.Service
}

func NewCmsNewsService() *CmsNewsService {
	return &CmsNewsService{
		&dzhCore.Service{
			Model: model.NewCmsNews(),
			ListQueryOp: &dzhCore.QueryOp{
				FieldEQ:      []string{},
				KeyWordField: []string{},
				AddOrderby:   map[string]string{"`createTime`": "DESC"},
			},
			PageQueryOp: &dzhCore.QueryOp{
				FieldEQ:      []string{},
				KeyWordField: []string{},
				AddOrderby:   map[string]string{"`createTime`": "DESC"},
			},
		},
	}
}
