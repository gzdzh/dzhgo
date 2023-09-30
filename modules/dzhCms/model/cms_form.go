package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameCmsForm = "dzh_cms_form"

// CmsForm mapped from table <cms_form>
type CmsForm struct {
	*dzhCore.Model
	Name   string `gorm:"column:name;not null;comment:姓名;type:varchar(10)" json:"name"`
	Phone  string `gorm:"column:phone;not null;comment:号码;type:varchar(11)" json:"phone"`
	Remark string `gorm:"column:remark;not null;comment:题目" json:"remark"`
}

// TableName CmsForm's table name
func (*CmsForm) TableName() string {
	return TableNameCmsForm
}

// GroupName CmsForm's table group
func (*CmsForm) GroupName() string {
	return "default"
}

// NewCmsForm create a new CmsForm
func NewCmsForm() *CmsForm {
	return &CmsForm{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&CmsForm{})
}
