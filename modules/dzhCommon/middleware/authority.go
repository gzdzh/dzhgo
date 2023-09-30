package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gzdzh/dzhgo/dzhCore"

	"github.com/gzdzh/dzhgo/modules/dzhCommon/config"
	// "github.com/gzdzh/dzhgo/modules/base/config"
)

// 本类接口无需权限验证
func CommonAuthorityMiddlewareOpen(r *ghttp.Request) {
	// g.Dump("CommonAuthorityMiddlewareOpen")
	r.SetCtxVar("AuthOpen", true)
	r.Middleware.Next()
}

// 本类接口无需权限验证,只需登录验证
func CommonAuthorityMiddlewareComm(r *ghttp.Request) {
	// g.Dump("CommonAuthorityMiddlewareComm")
	r.SetCtxVar("AuthComm", true)
	r.Middleware.Next()
}

func CommonAuthorityMiddleware(r *ghttp.Request) {

	var (
		statusCode = 200
		ctx        = r.GetCtx()
	)

	// 无需登录验证
	AuthOpen := r.GetCtxVar("AuthOpen", false)
	// fmt.Println("AuthOpen", AuthOpen)
	if AuthOpen.Bool() {
		r.Middleware.Next()
		return
	}
	// 	 解密token,拿到会员登录数据
	type Claims struct {
		IsRefresh       bool   `json:"isRefresh"`
		RoleIds         []uint `json:"roleIds"`
		Username        string `json:"username"`
		UserId          uint   `json:"userId"`
		NickName        string `json:"nickName"`
		PasswordVersion *int32 `json:"passwordVersion"`
		jwt.RegisteredClaims
	}

	tokenString := r.GetHeader("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.Config.Jwt.Secret), nil
	})
	if err != nil {
		g.Log().Error(ctx, "CommmonAuthorityMiddleware", err)
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    -1,
			"message": "登陆失效～",
		})
	}
	if !token.Valid {
		g.Log().Error(ctx, "CommmonAuthorityMiddleware", "token invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    -1,
			"message": "登陆失效～",
		})
	}
	member := token.Claims.(*Claims)
	// g.Dump("member", member)
	// 将用户信息放入上下文
	r.SetCtxVar("member", member)
	// log.Println("member", member)

	cachetoken, _ := dzhCore.CacheManager.Get(ctx, "member:token:"+gconv.String(member.UserId))
	rtoken := cachetoken.String()
	// fmt.Println("rtoken", rtoken)

	// 只验证登录不验证权限的接口
	AuthComm := r.GetCtxVar("AuthComm", false)
	// fmt.Println("AuthComm", AuthComm)
	if AuthComm.Bool() {
		r.Middleware.Next()

		return
	}
	// 	 拿到会员数据验证登录
	// 如果传的token是refreshToken则校验失败
	if member.IsRefresh {
		g.Log().Error(ctx, "CommmonAuthorityMiddleware", "token invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    -1,
			"message": "登陆失效～",
		})
	}

	// 如果rtoken为空
	if rtoken == "" {
		g.Log().Error(ctx, "CommmonAuthorityMiddleware", "rtoken invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    -1,
			"message": "登陆失效～",
		})
	}
	// 如果rtoken不等于token 且 sso 未开启
	if tokenString != rtoken && !config.Config.Jwt.Sso {
		g.Log().Error(ctx, "CommmonAuthorityMiddleware", "token invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    -1,
			"message": "登陆失效～",
		})
	}

	r.Middleware.Next()
}
