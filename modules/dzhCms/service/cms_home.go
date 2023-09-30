package service

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CmsHomeService struct {
	*dzhCore.Service
}

func (s *CmsHomeService) Home(ctx context.Context, req *v1.HomeReq) (data interface{}, err error) {
	var (
	// r      = g.RequestFromCtx(ctx)
	// reqmap = r.GetMap()
	)

	resultMap := g.MapStrAny{}

	// 新闻点击排序
	NewsList, _ := dzhCore.DBM(model.NewCmsNews()).OrderDesc("click").Limit(10).All()
	if !NewsList.IsEmpty() {
		resultMap["newsList"] = NewsList
	}

	// 首页幻灯片
	BannerList, _ := dzhCore.DBM(model.NewCmsBanner()).All()
	if !NewsList.IsEmpty() {
		resultMap["bannerList"] = BannerList
	}

	// 配置
	Setting, _ := dzhCore.DBM(model.NewCmsSetting()).Where("id=1").One()
	if !NewsList.IsEmpty() {
		resultMap["setting"] = Setting
	}

	data = resultMap
	return
}

func NewCmsHomeService() *CmsHomeService {
	return &CmsHomeService{
		&dzhCore.Service{},
	}
}
