package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameCmsBannerType = "dzh_cms_banner_type"

// CmsBannerType mapped from table <cms_banner_type>
type CmsBannerType struct {
	*dzhCore.Model
	Title    string  `gorm:"column:title;not null;comment:标题" json:"title"`
	Image    *string `gorm:"column:image;comment:图片" json:"image"`
	Abstract *string `gorm:"column:abstract;comment:摘要" json:"abstract"`
	Remark   *string `gorm:"column:remark;comment:内容" json:"remark"`
	Status   string  `gorm:"column:status;comment:状态;type:int" json:"status"`
	OrderNum int32   `gorm:"column:orderNum;type:int;not null;default:99" json:"orderNum"`
	ParentID uint    `gorm:"column:parentId;type:bigint" json:"parentId"`
}

// TableName CmsBannerType's table name
func (*CmsBannerType) TableName() string {
	return TableNameCmsBannerType
}

// GroupName CmsBannerType's table group
func (*CmsBannerType) GroupName() string {
	return "default"
}

// NewCmsBannerType create a new CmsBannerType
func NewCmsBannerType() *CmsBannerType {
	return &CmsBannerType{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&CmsBannerType{})
}
