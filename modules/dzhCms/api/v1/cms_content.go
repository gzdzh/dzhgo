package v1

import "github.com/gogf/gf/v2/frame/g"

// 接口请求
type IndexReq struct {
	g.Meta `path:"/index" method:"GET"`
}

type Index2Req struct {
	g.Meta `path:"/index2" method:"GET"`
}

type Index3Req struct {
	g.Meta `path:"/index3" method:"GET"`
}

type HomeReq struct {
	g.Meta `path:"/home" method:"POST"`
}

// 字典
type ProDataReq struct {
	g.Meta `path:"/proData" method:"POST"`
	Id     int64 `json:"id" v:"required"`
}

// 广告
type BannerListReq struct {
	g.Meta `path:"/bannerList" method:"POST"`
	TypeId int64 `json:"typeId"`
}

// 客服列表
type ContactListReq struct {
	g.Meta `path:"/contactList" method:"POST"`
}

// 专业详情
type SubjectInfoReq struct {
	g.Meta `path:"/subjectInfo" method:"GET"`
	Id     int64 `json:"id" v:"required"`
}
