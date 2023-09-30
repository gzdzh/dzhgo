package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameCmsSetting = "dzh_cms_setting"

// CmsSetting mapped from table <cms_setting>
type CmsSetting struct {
	*dzhCore.Model
	SiteName           string  `gorm:"column:sitename;comment:站点名称" json:"sitename"`
	Notic              *string `gorm:"column:notic;comment:公告" json:"notic"`
	Image              *string `gorm:"column:image;comment:图片" json:"image"`
	Contact            *string `gorm:"column:contact;comment:联系人;type:varchar(10)" json:"contact"`
	Mobile             *string `gorm:"column:mobile;comment:联系方式;type:varchar(11)" json:"mobile"`
	ContactList        *string `gorm:"column:contactList;comment:客服列表" json:"contactList"`
	Gpt4               int     `gorm:"column:gpt4;comment:gpt4;default:0" json:"gpt4"`
	BaiduTranApiKey    *string `gorm:"column:baiduTranApiKey;comment:百度翻译apikey" json:"baiduTranApiKey"`
	BaiduTranSecretKey *string `gorm:"column:baiduTranSecretKey;comment:百度翻译Secretkey" json:"baiduTranSecretKey"`
	GuildId            *string `gorm:"column:guild_id;comment:服务器Id" json:"guild_id"`
	ChannelId          *string `gorm:"column:channel_id;comment:频道Id" json:"channel_id"`
	BotToken           *string `gorm:"column:bot_token;comment:机器人" json:"bot_token"`
	CdnProxyUrl        *string `gorm:"column:cdn_proxy_url;comment:图片代理地址" json:"cdn_proxy_url"`
	DiscordToken       *string `gorm:"column:discordToken;comment:discordToken" json:"discordToken"`
	Phrase             *string `gorm:"column:phrase;comment:过滤词" json:"phrase"`
	Appid              *string `gorm:"column:appid;comment:普通商户appid" json:"appid"`
	MchId              *string `gorm:"column:mchId;comment:普通商户号" json:"mchId"`
	CAPIv3Key          *string `gorm:"column:cAPIv3Key;comment:收款商户v3密钥" json:"cAPIv3Key"`
	CSerialNo          *string `gorm:"column:cSerialNo;comment:序列号" json:"cSerialNo"`
	CNotifyUrl         *string `gorm:"column:cNotifyUrl;comment:支付回调地址" json:"cNotifyUrl"`
	SpMchid            *string `gorm:"column:spMchid;comment:服务商商户号" json:"spMchid"`
	SpAppid            *string `gorm:"column:spAppid;comment:服务商appid" json:"spAppid"`
	SubMchId           *string `gorm:"column:subMchId;comment:特约商户" json:"subMchId"`
	APIv3Key           *string `gorm:"column:aPIv3Key;comment:收款商户v3密钥" json:"aPIv3Key"`
	SerialNo           *string `gorm:"column:serialNo;comment:序列号" json:"serialNo"`
	NotifyUrl          *string `gorm:"column:notifyUrl;comment:支付回调地址" json:"notifyUrl"`
	PayType            *int    `gorm:"column:payType;comment:支付模式;not null;default:1;index" json:"payType"`
}

// TableName CmsSetting's table name
func (*CmsSetting) TableName() string {
	return TableNameCmsSetting
}

// GroupName CmsSetting's table group
func (*CmsSetting) GroupName() string {
	return "default"
}

// NewCmsSetting create a new CmsSetting
func NewCmsSetting() *CmsSetting {
	return &CmsSetting{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&CmsSetting{})
}
