package v1

import "github.com/gogf/gf/v2/frame/g"

// 接口请求
type FormAddReq struct {
	g.Meta `path:"/formAdd" method:"POST"`
	Name   int    `json:"name"`
	Phone  int    `json:"phone"`
	Remark string `json:"remark"`
}
