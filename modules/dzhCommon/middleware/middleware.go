package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

func init() {

	s := g.Server()

	// 路由标记
	s.BindMiddleware("/app/cms/*", CommonAuthorityMiddlewareOpen)
	s.BindMiddleware("/app/*/open/*", CommonAuthorityMiddlewareOpen)
	s.BindMiddleware("/app/*/comm/*", CommonAuthorityMiddlewareComm)

	// 开始鉴权
	s.BindMiddleware("/app/*", CommonAuthorityMiddleware)

	// 静态目录映射
	// 上传文件
	if gfile.Exists(gfile.MainPkgPath() + "/public") {
		s.AddStaticPath("/dzh/public", gfile.MainPkgPath()+"/public")
	}

	// vue静态文件
	if gfile.Exists(gfile.MainPkgPath() + "/template/assets") {
		s.AddStaticPath("/assets", gfile.MainPkgPath()+"/template/assets")
	}

	// 普通静态文件
	if gfile.Exists(gfile.MainPkgPath() + "/template") {
		s.AddStaticPath("/template", gfile.MainPkgPath()+"/template")
	}

}
