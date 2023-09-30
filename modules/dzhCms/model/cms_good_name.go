package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameCmsGoodName = "dzh_cms_good_name"

// CmsGoodName mapped from table <cms_good_name>
type CmsGoodName struct {
	*dzhCore.Model
	Name     string  `gorm:"column:name;not null;comment:名称" json:"name"`
	Img      *string `gorm:"column:img;comment:图片" json:"img"`
	OrderNum int     `gorm:"column:orderNum;comment:排序;default:99" json:"orderNum"`
	Status   *int32  `gorm:"column:status;comment:状态" json:"status"`
	Remark   *string `gorm:"column:remark;comment:备注" json:"remark"`
	TypeId   int64   `gorm:"column:typeId;not null;comment:类别;index" json:"typeId"`
}

// TableName CmsGoodName's table name
func (*CmsGoodName) TableName() string {
	return TableNameCmsGoodName
}

// GroupName CmsGoodName's table group
func (*CmsGoodName) GroupName() string {
	return "default"
}

// NewCmsGoodName create a new CmsGoodName
func NewCmsGoodName() *CmsGoodName {
	return &CmsGoodName{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&CmsGoodName{})
}
