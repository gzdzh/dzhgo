package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameCommonSetting = "dzh_common_setting"

// CommonSetting mapped from table <common_setting>
type CommonSetting struct {
	*dzhCore.Model
	SiteName        *string `gorm:"column:siteName;comment:网站名称;varchar(50)" json:"siteName"`
	DomainName      *string `gorm:"column:domainName;comment:网站域名;type:varchar(50)" json:"domainName"`
	Logo            *string `gorm:"column:logo;comment:logo;type:varchar(50)" json:"logo"`
	Company         *string `gorm:"column:company;comment:公司名称;type:varchar(50)" json:"company"`
	Contact         *string `gorm:"column:contact;comment:联系人;type:varchar(50)" json:"contact"`
	ContactWay      *string `gorm:"column:contactWay;comment:座机;type:varchar(50)" json:"contactWay"`
	Mobile          *string `gorm:"column:mobile;comment:手机;type:varchar(50)" json:"mobile"`
	Address         *string `gorm:"column:Address;comment:地址;type:varchar(50)" json:"address"`
	Keyword         *string `gorm:"column:keyword;comment:关键词;type:varchar(50)" json:"keyword"`
	Description     *string `gorm:"column:description;comment:描述;type:varchar(50)" json:"description"`
	Smtp            *string `gorm:"column:smtp;comment:smtp;type:varchar(50)" json:"smtp"`
	SendEmail       *string `gorm:"column:sendEmail;comment:发送邮箱;type:varchar(50)" json:"sendEmail"`
	Pass            *string `gorm:"column:pass;comment:邮箱授权码;type:varchar(50)" json:"pass"`
	RequestEmail    *string `gorm:"column:requestEmail;comment:接收邮箱;type:varchar(50)" json:"requestEmail"`
	RemindEmail     *int    `gorm:"column:remindEmail;comment:到期邮件开启 0关闭 1开启;default:0;type:int(11)" json:"remindEmail"`
	RemindSms       *int    `gorm:"column:remindSms;comment:到期短信开启 0关闭 1开启;default:0;type:int(11)" json:"remindSms"`
	AccessKeyId     *string `gorm:"column:accessKeyId;comment:accessKeyId;type:varchar(100)" json:"accessKeyId"`
	AccessKeySecret *string `gorm:"column:accessKeySecret;comment:accessKeySecret;type:varchar(100)" json:"accessKeySecret"`
	SignName        *string `gorm:"column:signName;comment:签名;type:varchar(50)" json:"signName"`
	TemplateCode    *string `gorm:"column:templateCode;comment:模板;type:varchar(50)" json:"templateCode"`
	Endpoint        *string `gorm:"column:endpoint;comment:endpoint;type:varchar(50)" json:"endpoint"`
	RemindMobile    *string `gorm:"column:remindMobile;comment:通知手机号码;type:varchar(50)" json:"remindMobile"`
	RemindDay       *string `gorm:"column:remindDay;comment:到期提醒提前天数;type:varchar(50)" json:"remindDay"`
	FieldJson       *string `gorm:"column:fieldJson;comment:自定义字段" json:"fieldJson"`
}

// TableName CommonSetting's table name
func (*CommonSetting) TableName() string {
	return TableNameCommonSetting
}

// GroupName CommonSetting's table group
func (*CommonSetting) GroupName() string {
	return "default"
}

// NewCommonSetting create a new CommonSetting
func NewCommonSetting() *CommonSetting {
	return &CommonSetting{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&CommonSetting{})
}
