package service

import (
	"context"
	"fmt"

	"github.com/gzdzh/dzhgo/dzhCore"
	v1 "github.com/gzdzh/dzhgo/modules/dzhCms/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhCms/model"
	commonM "github.com/gzdzh/dzhgo/modules/dzhCommon/model"
	"github.com/gzdzh/dzhgo/modules/dzhCommon/service"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type CmsFormService struct {
	*dzhCore.Service
}

func (s *CmsFormService) FormAdd(ctx context.Context, req *v1.FormAddReq) (data interface{}, err error) {
	var (
		m      = dzhCore.DBM(s.Model)
		r      = g.RequestFromCtx(ctx)
		reqmap = r.GetMap()
	)

	var formInsert gmap.AnyAnyMap
	formInsert.Set("name", reqmap["name"])
	formInsert.Set("phone", reqmap["phone"])

	var str string
	if remark, ok := reqmap["remark"].([]interface{}); ok {
		for _, value := range remark {
			if v, ok := value.(map[string]interface{}); ok {
				str += gconv.String(v["num"]) + "、" + gconv.String(v["title"]) + "(" + gconv.String(v["value"]) + ")" + gconv.String(v["unit"]) + "</br>"
			}
		}
	}

	formInsert.Set("remark", str)

	config, err := dzhCore.DBM(commonM.NewCommonSetting()).Where("id=1").One()
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		err = gerror.New("读取配置失败")
	}
	if result, err := m.Data(formInsert.Map()).InsertAndGetId(); err == nil {
		fmt.Println("留言")
		data = result
		content := fmt.Sprintf("请打开后台，你有新的留言,下单信息，姓名：%s，电话：%s", reqmap["name"], reqmap["phone"])
		subject := "有新的留言"
		addressHeader := "通知邮件"
		service.NewCommonSentService().SentEmail(content, subject, addressHeader, config)
	}

	return
}

func NewCmsFormService() *CmsFormService {
	return &CmsFormService{
		&dzhCore.Service{
			Model:       model.NewCmsForm(),
			ListQueryOp: &dzhCore.QueryOp{},
			PageQueryOp: &dzhCore.QueryOp{
				FieldEQ:      []string{},
				KeyWordField: []string{},
				AddOrderby:   map[string]string{"createTime": "DESC"},
			},
			InfoIgnoreProperty: "",
			UniqueKey:          map[string]string{},
			NotNullKey:         map[string]string{},
		},
	}
}
