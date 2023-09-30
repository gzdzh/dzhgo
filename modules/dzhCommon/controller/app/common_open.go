package app

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"
)

type DzhcommonOpenController struct {
	*dzhCore.ControllerSimple
}

func init() {
	var dzhcommon_open_controller = &DzhcommonOpenController{
		&dzhCore.ControllerSimple{
			Perfix: "/app/common/open",
		},
	}
	// 注册路由
	dzhCore.RegisterControllerSimple(dzhcommon_open_controller)
}

type DzhcommonOpenWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type DzhcommonOpenWelcomeRes struct {
	*dzhCore.BaseRes
	Data interface{} `json:"data"`
}

func (c *DzhcommonOpenController) Welcome(ctx context.Context, req *DzhcommonOpenWelcomeReq) (res *DzhcommonOpenWelcomeRes, err error) {
	res = &DzhcommonOpenWelcomeRes{
		BaseRes: dzhCore.Ok("Welcome /app/common/open"),
		Data:    gjson.New(`{"name": "/app/common/open", "age":0}`),
	}
	return
}

type RequestReq struct {
	g.Meta   `path:"/requestData" method:"GET"`
	UserName string `json:"usernmae"`
	NickName string `json:"nickname"`
}
type RequestRes struct {
	*dzhCore.BaseRes
	Data interface{} `json:"data"`
}

func (c *DzhcommonOpenController) RequestData(ctx context.Context, req *RequestReq) (res *RequestRes, err error) {
	r := g.RequestFromCtx(ctx)
	rmap := r.GetMap()
	g.Dump("rmap", rmap)

	g.Dump("req", req)
	res = &RequestRes{
		BaseRes: dzhCore.Ok("/app/common/open/requestData"),
		Data:    req,
	}
	return
}

type BaseCommUploadReq struct {
	g.Meta        `path:"/upload" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

// Upload 方法
func (c *DzhcommonOpenController) Upload(ctx context.Context, req *BaseCommUploadReq) (res *dzhCore.BaseRes, err error) {
	data, err := dzhCore.File().Upload(ctx)
	res = dzhCore.Ok(data)
	return
}
