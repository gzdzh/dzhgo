package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"
)

type CmsGoodNameService struct {
	*dzhCore.Service
}

func NewCmsGoodNameService() *CmsGoodNameService {
	return &CmsGoodNameService{
		&dzhCore.Service{
			Model: model.NewCmsGoodName(),
			PageQueryOp: &dzhCore.QueryOp{
				FieldEQ:      []string{},
				KeyWordField: []string{"title"},
				AddOrderby:   map[string]string{"`createTime`": "DESC"},
				Where: func(ctx context.Context) [][]interface{} {
					r := g.RequestFromCtx(ctx).GetMap()
					return [][]interface{}{{"typeId in (?)", g.Slice{r["typeIds"]}, r["typeIds"]}}
				},
			},
		},
	}
}
