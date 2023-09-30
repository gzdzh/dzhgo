package model

import "github.com/gzdzh/dzhgo/dzhCore"

const TableNameBaseSysRoleDepartment = "base_sys_role_department"

// BaseSysRoleDepartment mapped from table <base_sys_role_department>
type BaseSysRoleDepartment struct {
	*dzhCore.Model
	RoleID       uint `gorm:"column:roleId;type:bigint;not null" json:"roleId"`             // 角色ID
	DepartmentID uint `gorm:"column:departmentId;type:bigint;not null" json:"departmentId"` // 部门ID
}

// TableName BaseSysRoleDepartment's table name
func (*BaseSysRoleDepartment) TableName() string {
	return TableNameBaseSysRoleDepartment
}

// NewBaseSysRoleDepartment create a new BaseSysRoleDepartment
func NewBaseSysRoleDepartment() *BaseSysRoleDepartment {
	return &BaseSysRoleDepartment{
		Model: dzhCore.NewModel(),
	}
}

// init 创建表
func init() {
	dzhCore.CreateTable(&BaseSysRoleDepartment{})
}
