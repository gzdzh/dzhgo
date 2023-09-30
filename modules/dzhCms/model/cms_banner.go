package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameCmsBanner = "dzh_cms_banner"

// CmsBanner mapped from table <cms_banner>
type CmsBanner struct {
	*dzhCore.Model
	Title    string  `gorm:"column:title;not null;comment:标题" json:"title"`
	Image    *string `gorm:"column:image;comment:图片" json:"image"`
	Link     *string `gorm:"column:link;comment:跳转" json:"link"`
	TypeId   int64   `gorm:"column:typeId;comment:类别;index" json:"typeId"`
	Remark   *string `gorm:"column:remark;comment:内容" json:"remark"`
	Status   string  `gorm:"column:status;comment:状态;type:int" json:"status"`
	OrderNum int32   `gorm:"column:orderNum;comment:排序;type:int;not null;default:99" json:"orderNum"`
}

// TableName CmsBanner's table name
func (*CmsBanner) TableName() string {
	return TableNameCmsBanner
}

// GroupName CmsBanner's table group
func (*CmsBanner) GroupName() string {
	return "default"
}

// NewCmsBanner create a new CmsBanner
func NewCmsBanner() *CmsBanner {
	return &CmsBanner{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&CmsBanner{})
}
