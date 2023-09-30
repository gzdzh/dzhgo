package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameCmsGoodType = "dzh_cms_good_type"

// CmsGoodType mapped from table <cms_good_type>
type CmsGoodType struct {
	*dzhCore.Model
	Name     string  `gorm:"column:name;not null;comment:名称" json:"name"`
	Img      *string `gorm:"column:img;comment:图片" json:"img"`
	OrderNum int     `gorm:"column:orderNum;comment:排序;default:99" json:"orderNum"`
	Status   *int32  `gorm:"column:status;comment:状态" json:"status"`
	Remark   *string `gorm:"column:remark;comment:备注" json:"remark"`
}

// TableName CmsGoodType's table name
func (*CmsGoodType) TableName() string {
	return TableNameCmsGoodType
}

// GroupName CmsGoodType's table group
func (*CmsGoodType) GroupName() string {
	return "default"
}

// NewCmsGoodType create a new CmsGoodType
func NewCmsGoodType() *CmsGoodType {
	return &CmsGoodType{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&CmsGoodType{})
}
