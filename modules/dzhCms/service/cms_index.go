package service

import (
	"context"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"
)

type CmsIndexService struct {
	*dzhCore.Service
}

// 发送邮件
func (s *CmsIndexService) Index(ctx context.Context, req *v1.IndexReq) (result *interface{}, err error) {

	return
}

func NewCmsIndexService() *CmsIndexService {
	return &CmsIndexService{
		&dzhCore.Service{},
	}
}
