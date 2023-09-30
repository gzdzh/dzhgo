package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/modules/dzhFinance/model"
)

// 普通商户native下单
type WxV3NativeReq struct {
	g.Meta `path:"/wxV3Native" method:"POST"`
}

// 服务商native下单
type V3PartnerNativeReq struct {
	g.Meta `path:"/V3PartnerNative" method:"POST"`
	*model.PayModel
}

type Resource struct {
	OriginalType   string `json:"original_type,omitempty"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type SignInfo struct {
	HeaderTimestamp string `json:"Wechatpay-Timestamp"`
	HeaderNonce     string `json:"Wechatpay-Nonce"`
	HeaderSignature string `json:"Wechatpay-Signature"`
	HeaderSerial    string `json:"Wechatpay-Serial"`
	SignBody        string `json:"sign_body"`
}

type V3NotifyReq struct {
	Id           string    `json:"id"`
	CreateTime   string    `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     *Resource `json:"resource"`
	SignInfo     *SignInfo `json:"-"`
}

// 支付后回调
type PayNoticeReq struct {
	g.Meta `path:"/payNotice" method:"POST"`
	V3NotifyReq
}

// 支付后回调
type NoticeReq struct {
	g.Meta `path:"/notice" method:"POST"`
	V3NotifyReq
}
