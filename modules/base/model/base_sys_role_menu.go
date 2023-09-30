package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameBaseSysRoleMenu = "base_sys_role_menu"

// BaseSysRoleMenu mapped from table <base_sys_role_menu>
type BaseSysRoleMenu struct {
	*dzhCore.Model
	RoleID uint `gorm:"column:roleId;type:bigint;not null" json:"roleId"` // 角色ID
	MenuID uint `gorm:"column:menuId;type:bigint;not null" json:"menuId"` // 菜单ID
}

// TableName BaseSysRoleMenu's table name
func (*BaseSysRoleMenu) TableName() string {
	return TableNameBaseSysRoleMenu
}

// NewBaseSysRoleMenu create a new BaseSysRoleMenu
func NewBaseSysRoleMenu() *BaseSysRoleMenu {
	return &BaseSysRoleMenu{
		Model: &dzhCore.Model{},
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&BaseSysRoleMenu{})
}
