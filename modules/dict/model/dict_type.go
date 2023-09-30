package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameDictType = "dict_type"

// DictType mapped from table <dict_type>
type DictType struct {
	*dzhCore.Model
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"` // 名称
	Key  string `gorm:"column:key;type:varchar(255);not null" json:"key"`   // 标识
}

// TableName DictType's table name
func (*DictType) TableName() string {
	return TableNameDictType
}

// GroupName DictType's table group
func (*DictType) GroupName() string {
	return "default"
}

// NewDictType create a new DictType
func NewDictType() *DictType {
	return &DictType{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&DictType{})
}
