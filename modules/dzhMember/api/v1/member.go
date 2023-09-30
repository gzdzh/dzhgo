package v1

import "github.com/gogf/gf/v2/frame/g"

// 更新会员信息
type UpdatePersonReq struct {
	g.Meta `path:"/updatePerson" method:"POST"`
}

type DecrementMJCountReq struct {
	g.Meta `path:"/decrementMJCount" method:"POST"`
}

type MemberInfoReq struct {
	g.Meta `path:"/memberInfo" method:"POST"`
}
