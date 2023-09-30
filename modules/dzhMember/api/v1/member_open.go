package v1

import "github.com/gogf/gf/v2/frame/g"

// login 接口请求
type LoginReq struct {
	g.Meta    `path:"/login" method:"POST"`
	JsCode    string `json:"code" p:"code" v:"required"`
	NickName  string `json:"nickName" p:"nickName" v:"required"`
	AvatarUrl string `json:"avatarUrl" p:"avatarUrl" v:"required"`
	LevelName string `json:"levelName" p:"levelName" v:"levelName"`
}

type MpLoginReq struct {
	g.Meta `path:"/mpLogin" method:"POST"`
}

// WxLogin 接口请求 微信登录
type WxLoginReq struct {
	g.Meta    `path:"/wx" method:"POST"`
	JsCode    string `json:"code" p:"code" v:"required"`
	UserName  string `json:"userName" p:"userName" v:"required"`
	NickName  string `json:"nickName" p:"nickName" v:"required"`
	AvatarUrl string `json:"avatarUrl" p:"avatarUrl" v:"required"`
	LevelName string `json:"levelName" p:"levelName" v:"required"`
}

// 游客登录
type TouristLoginReq struct {
	g.Meta `path:"/tourist" method:"POST"`
}

// 验证游客次数
type VerifyCountReq struct {
	g.Meta `path:"/verifyCount" method:"POST"`
	Token  string `json:"token" p:"token" v:"required"`
}

// 账号登录
type AccountLoginReq struct {
	g.Meta   `path:"/account" method:"POST"`
	UserName string `json:"userName" p:"username" v:"required"`
	PassWord string `json:"passWord" p:"password" v:"required"`
}

// 账号注册
type AccountRegisterReq struct {
	g.Meta     `path:"/accountReg" method:"POST"`
	UserName   string `json:"userName" p:"userName" v:"required"`
	PassWord   string `json:"passWord" p:"passWord" v:"required"`
	RePassWord string `json:"rePassWord" p:"rePassWord" v:"required"`
}
