package middleware

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

func IndexMiddleware(r *ghttp.Request) {
	indexHtm := gfile.Exists(gfile.MainPkgPath() + "/template/index.htm")
	indexHtml := gfile.Exists(gfile.MainPkgPath() + "/template/index.html")
	indexTpl := gfile.Exists(gfile.MainPkgPath() + "/template/index.tpl")
	var templateName string
	if indexHtml {
		templateName = "index.html"
		r.SetParam("templateName", templateName)
		r.Middleware.Next()
		return
	}

	if indexHtm {
		templateName = "index.htm"
		r.SetParam("templateName", templateName)
		r.Middleware.Next()
		return
	}

	if indexTpl {
		templateName = "index.tpl"
		r.SetParam("templateName", templateName)
		r.Middleware.Next()
		return
	}

}

func init() {
	var ctx g.Ctx
	version, _ := g.Cfg().Get(ctx, "dzhcms.version")
	fmt.Printf("dzhcms version: %v\n", version)
	s := g.Server()
	// 预处理首页模版
	s.BindMiddleware("/app/cms/index/index", IndexMiddleware)

	// s.BindHandler("/", func(r *ghttp.Request) {
	// 	r.Response.WriteTpl(r.GetParam("templateName").String(), g.Map{
	// 		"name": "姓名",
	// 	})
	// })

	g.Dump("cms")
}
