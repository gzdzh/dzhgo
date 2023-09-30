package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gzdzh/dzhgo/dzhCore"

	v1 "github.com/gzdzh/dzhgo/modules/dzhMember/api/v1"
	"github.com/gzdzh/dzhgo/modules/dzhMember/config"
	"github.com/gzdzh/dzhgo/modules/dzhMember/model"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/golang-jwt/jwt/v4"

	commonCof "github.com/gzdzh/dzhgo/modules/dzhCommon/config"
)

type MemberUserService struct {
	*dzhCore.Service
}

type TokenResult struct {
	Expire         uint                   `json:"expire"`
	Token          string                 `json:"token"`
	RefreshExpires uint                   `json:"refreshExpires"`
	RefreshToken   string                 `json:"refreshToken"`
	UserInfo       map[string]interface{} `json:"userInfo"`
}

type LoginData struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
}
type LoginReq struct {
	g.Meta    `path:"/login" method:"POST"`
	JsCode    string `json:"code" p:"code" v:"required"`
	NickName  string `json:"nickName" p:"nickName" v:"required"`
	AvatarUrl string `json:"avatarUrl" p:"avatarUrl" v:"required"`
	LevelName string `json:"levelName" p:"levelName" v:"levelName"`
}

func (m *MemberUserService) Login(ctx g.Ctx, req *v1.LoginReq) (result *TokenResult, err error) {

	url := "https://api.weixin.qq.com/sns/jscode2session"
	params := g.Map{
		"appid":      "wx44b19e138b892f99",
		"secret":     "1e483dec348f5749f5dc30e3ec6d0afa",
		"js_code":    req.JsCode,
		"grant_type": "authorization_code",
	}

	// 请求腾讯开发者服务器,code换取openid
	codeRes, err := g.Client().Post(ctx, url, params)
	if err != nil {
		panic(err)
	}
	defer codeRes.Close()

	// 读取响应体并且json转换到结构体
	loginData := &LoginData{}
	err = json.NewDecoder(codeRes.Body).Decode(loginData)
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		err = gerror.New("读取响应体失败")
		return
	}
	var (
		member     *model.MemberUser
		memberId   interface{}
		memberData g.Map
	)

	dzhCore.DBM(m.Model).Where("openid=?", loginData.Openid).Where("status=?", 1).Scan(&member)
	if member == nil {
		md5password, _ := gmd5.Encrypt("123456")
		memberData = g.Map{
			"username":   req.NickName,
			"password":   md5password,
			"nickName":   req.NickName,
			"levelName":  "至尊VIP",
			"avatarUrl":  req.AvatarUrl,
			"openid":     loginData.Openid,
			"unionid":    loginData.Unionid,
			"sessionKey": loginData.SessionKey,
		}
		memberId, err = dzhCore.DBM(m.Model).Data(memberData).InsertAndGetId()
		if err != nil {
			g.Log().Error(ctx, "func Login insert :"+err.Error(), err)
			err = gerror.New("读取响应体失败")
			return
		}
		memberData["id"] = memberId
		gconv.Struct(memberData, &member)
	} else {
		memberData = g.Map{
			"username":  req.NickName,
			"avatarUrl": req.AvatarUrl,
		}
		_, err = dzhCore.DBM(m.Model).Where("openid=?", loginData.Openid).Data(memberData).Update()
		err = dzhCore.DBM(m.Model).Where("openid=?", loginData.Openid).Scan(&member)
		if err != nil {
			err = gerror.New("查询会员失败")
			return
		}
	}

	result, err = m.generateTokenByUser(ctx, member)
	if err != nil {
		return
	}

	return
}

// 会员信息
func (m *MemberUserService) MemberInfo(ctx g.Ctx) (data interface{}, err error) {
	var (
		r = g.RequestFromCtx(ctx)
	)
	member := r.GetCtxVar("member")

	data = member
	return
}

// 微信公众号登录
func (m *MemberUserService) MpLoginReq(ctx g.Ctx, req *v1.MpLoginReq) (result *TokenResult, err error) {

	var (
		r    = g.RequestFromCtx(ctx)
		rmap = r.GetMap()
	)

	_, err = dzhCore.DBM(m.Model).Data(rmap).Save()
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		err = gerror.New("录入失败")
		return
	}

	fmt.Printf("rmap: %v\n", rmap)

	return
}

// 微信登录
func (m *MemberUserService) WxLogin(ctx g.Ctx, req *v1.WxLoginReq) (result *TokenResult, err error) {

	wxCon := config.GetWxCon()
	url := gconv.String(wxCon["requestUrl"])
	params := g.Map{
		"appid":      gconv.String(wxCon["appid"]),
		"secret":     gconv.String(wxCon["secret"]),
		"js_code":    req.JsCode,
		"grant_type": gconv.String(wxCon["grantType"]),
	}

	// 请求腾讯开发者服务器,code换取openid
	codeRes, err := g.Client().Post(ctx, url, params)
	if err != nil {
		panic(err)
	}
	if gconv.String(codeRes.Body) == "{}" {
		err = gerror.New("请求微信服务器错误")
		g.Log().Error(ctx, err.Error(), err)
		return
	}
	defer codeRes.Close()

	// 读取响应体并且json转换到结构体
	loginData := &LoginData{}
	err = json.NewDecoder(codeRes.Body).Decode(loginData)
	if err != nil {
		g.Log().Error(ctx, err.Error(), err)
		err = gerror.New("读取响应体失败")
		return
	}
	var (
		member     *model.MemberUser
		memberId   interface{}
		memberData g.Map
	)

	dzhCore.DBM(m.Model).Where("openid=?", loginData.Openid).Where("status=?", 1).Scan(&member)
	// 不存在
	if member == nil {
		md5password, _ := gmd5.Encrypt("123456")
		memberData = g.Map{
			"username":   req.NickName,
			"password":   md5password,
			"nickName":   req.NickName,
			"levelName":  "至尊VIP",
			"avatarUrl":  req.AvatarUrl,
			"openid":     loginData.Openid,
			"unionid":    loginData.Unionid,
			"sessionKey": loginData.SessionKey,
		}
		memberId, err = dzhCore.DBM(m.Model).Data(memberData).InsertAndGetId()
		if err != nil {
			g.Log().Error(ctx, "func Login insert :"+err.Error(), err)
			err = gerror.New("读取响应体失败")
			return
		}
		memberData["id"] = memberId
		gconv.Struct(memberData, &member)
	} else {
		memberData = g.Map{
			"username":  req.NickName,
			"avatarUrl": req.AvatarUrl,
		}
		_, err = dzhCore.DBM(m.Model).Where("openid=?", loginData.Openid).Data(memberData).Update()
		err = dzhCore.DBM(m.Model).Where("openid=?", loginData.Openid).Scan(&member)
		if err != nil {
			err = gerror.New("查询会员失败")
			return
		}
	}

	result, err = m.generateTokenByUser(ctx, member)
	if err != nil {
		return
	}

	return
}

// 游客登录
func (m *MemberUserService) TouristLogin(ctx g.Ctx, req *v1.TouristLoginReq) (result *TokenResult, err error) {

	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	token := grand.Str(characters, 20)

	// 获取用户角色
	roleIds := make([]uint, 1)
	roleIds[0] = 1

	result = &TokenResult{
		UserInfo: g.Map{
			"avatarUrl":   "",
			"name":        "游客",
			"description": "购买套餐，享受更多优惠",
			"baseURI":     "http://www.gzdzh.cn",
			"accessToken": token,
			"id":          "",
		},
	}
	result.Token = token

	// 将用户相关信息保存到缓存
	dzhCore.CacheManager.Set(ctx, "member:token:"+gconv.String(token), token, 0)
	dzhCore.CacheManager.Set(ctx, "member:token:count:"+gconv.String(token), 0, 0)
	dzhCore.CacheManager.Set(ctx, "member:token:tourist:"+gconv.String(token), "true", 0)

	return
}

// 验证游客次数
func (m *MemberUserService) VerifyCount(ctx g.Ctx, req *v1.VerifyCountReq) (data interface{}, err error) {

	if req.Token == "" {
		err = gerror.New("token不能为空")
		g.Log().Error(ctx, "func VerifyCount :"+err.Error(), err)
		return
	}

	token, _ := dzhCore.CacheManager.Get(ctx, "member:token:"+gconv.String(req.Token))
	if token.IsEmpty() {
		err = gerror.New("账号过期，请重新登陆")
		g.Log().Error(ctx, "func VerifyCount :"+err.Error(), err)
		return
	}
	tokenCount, _ := dzhCore.CacheManager.Get(ctx, "member:token:count:"+gconv.String(req.Token))

	fmt.Printf("tokenCount: %v\n", tokenCount)
	if tokenCount.Int() >= 2 {
		err = gerror.New("超过体验次数，请购买套餐")
		g.Log().Error(ctx, "func VerifyCount :"+err.Error(), err)
		return
	}
	count := tokenCount.Int() + 1
	dzhCore.CacheManager.Set(ctx, "member:token:count:"+gconv.String(token), count, 0)

	return
}

// 账号登录
func (m *MemberUserService) AccountLogin(ctx g.Ctx, req *v1.AccountLoginReq) (result *TokenResult, err error) {

	var member *model.MemberUser
	pw, _ := gmd5.Encrypt(req.PassWord)
	dzhCore.DBM(m.Model).Where("userName=?", req.UserName).Where("passWord=?", pw).Scan(&member)
	if member == nil {
		err = gerror.New("账号密码错误")
		g.Log().Error(ctx, "func AccountLogin find :"+err.Error(), err)
		return
	}

	// 生成token
	result, err = m.generateTokenByUser(ctx, member)
	if err != nil {
		return
	}

	return
}

// 账号注册
func (m *MemberUserService) AccountRegister(ctx g.Ctx, req *v1.AccountRegisterReq) (result *TokenResult, err error) {

	var (
		r    = g.RequestFromCtx(ctx)
		rmap = r.GetMap()
	)

	var member *model.MemberUser
	dzhCore.DBM(m.Model).Where("userName=?", req.UserName).Scan(&member)
	if member != nil {
		err = gerror.New("账号已存在")
		g.Log().Error(ctx, "func AccountRegister exit :"+err.Error(), err)
		return
	}

	if rmap["passWord"] != rmap["rePassWord"] {
		err = gerror.New("两次密码不一致！")
		g.Log().Error(ctx, "AccountRegister:"+err.Error(), err)
		return
	}

	pwMd5, _ := gmd5.Encrypt(req.PassWord)
	insertData := g.Map{
		"userName": req.UserName,
		"passWord": pwMd5,
	}
	_, err = dzhCore.DBM(m.Model).Data(insertData).Insert()
	if err != nil {
		g.Log().Error(ctx, "AccountRegister insert:"+err.Error(), err)
		return
	}

	return
}

// 根据用户生成前端需要的Token信息
func (m *MemberUserService) generateTokenByUser(ctx g.Ctx, member *model.MemberUser) (result *TokenResult, err error) {
	// 获取用户角色
	roleIds := make([]uint, 1)
	roleIds[0] = 1
	// 生成token
	result = &TokenResult{
		UserInfo: g.Map{
			"levelName": member.LevelName,
			"id":        member.ID,
		},
	}
	result.Expire = commonCof.Config.Jwt.Token.Expire
	result.RefreshExpires = commonCof.Config.Jwt.Token.RefreshExpire
	result.Token = m.generateToken(ctx, member, roleIds, result.Expire, false)
	result.RefreshToken = m.generateToken(ctx, member, roleIds, result.RefreshExpires, true)
	// 将用户相关信息保存到缓存

	dzhCore.CacheManager.Set(ctx, "member:token:"+gconv.String(member.ID), result.Token, 0)
	dzhCore.CacheManager.Set(ctx, "member:token:refresh:"+gconv.String(member.ID), result.RefreshToken, 0)

	return
}

// generateToken  生成token
func (m *MemberUserService) generateToken(ctx g.Ctx, member *model.MemberUser, roleIds []uint, exprire uint, isRefresh bool) (token string) {
	err := dzhCore.CacheManager.Set(ctx, "member:passwordVersion:"+gconv.String(member.ID), gconv.String(member.PasswordV), 0)
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}
	type Claims struct {
		IsRefresh       bool   `json:"isRefresh"`
		RoleIds         []uint `json:"roleIds"`
		Username        string `json:"username"`
		NickName        string `json:"nickName"`
		UserId          uint   `json:"userId"`
		PasswordVersion *int32 `json:"passwordVersion"`
		jwt.RegisteredClaims
	}
	claims := &Claims{
		IsRefresh:       isRefresh,
		RoleIds:         roleIds,
		Username:        member.Username,
		NickName:        member.NickName,
		UserId:          member.ID,
		PasswordVersion: member.PasswordV,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exprire) * time.Second)),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString([]byte(commonCof.Config.Jwt.Secret))
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}
	return
}

// ModifyBefore 新增|删除|修改前的操作
func (m *MemberUserService) ModifyBefore(ctx g.Ctx, method string, param g.MapStrAny) (err error) {

	r := g.RequestFromCtx(ctx)
	if method == "Add" || method == "Update" {
		password := r.Get("password").String()
		if password != "" {
			password = gmd5.MustEncryptString(password)
			r.SetParam("password", password)
		}
	}
	return
}

// ModifyAfter 新增|删除|修改后的操作
func (m *MemberUserService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	return
}

func NewMemberUserService() *MemberUserService {
	return &MemberUserService{
		&dzhCore.Service{
			Model:              model.NewMemberUser(),
			InfoIgnoreProperty: "password,passwordV",
		},
	}
}
