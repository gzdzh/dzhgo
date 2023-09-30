package service

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type CmsBannerService struct {
	*dzhCore.Service
}

// 图片集
func (s *CmsBannerService) BannerList(ctx context.Context, req *v1.BannerListReq) (data interface{}, err error) {

	var (
		bannerWhere = g.Map{}
		typeWhere   = g.Map{}
	)
	if req.TypeId != 0 {
		bannerWhere["typeId"] = req.TypeId
		typeWhere["id"] = req.TypeId

	}

	bannerList, _ := dzhCore.DBM(s.Model).Where(bannerWhere).All()
	bannerType, _ := dzhCore.DBM(model.NewCmsBannerType()).Where(typeWhere).All()

	result := g.Slice{} // 用于存储最终的结果
	for _, itemType := range bannerType {
		typeName := gconv.String(itemType["title"])
		children := g.Slice{}
		for _, bannerItem := range bannerList {
			if gconv.Int(itemType["id"]) == gconv.Int(bannerItem["typeId"]) {
				children = append(children, g.MapStrAny{
					"title": gconv.String(bannerItem["title"]),
					"image": gconv.String(bannerItem["image"]),
					"link":  gconv.String(bannerItem["link"]),
				})
			}
		}

		typeMap := g.MapStrAny{
			"typeName": typeName,
			"children": children,
		}
		result = append(result, typeMap)
	}

	data = result

	return
}

func NewCmsBannerService() *CmsBannerService {
	return &CmsBannerService{
		&dzhCore.Service{
			Model: model.NewCmsBanner(),
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
