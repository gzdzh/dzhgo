package service

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/model"
)

type BaseSysDepartmentService struct {
	*dzhCore.Service
}

// GetByRoleIds 获取部门
func (s *BaseSysDepartmentService) GetByRoleIds(roleIds []string, isAdmin bool) (res []uint) {
	var (
		result                gdb.Result
		BaseSysRoleDepartment = model.NewBaseSysRoleDepartment()
	)
	// 如果roleIds不为空
	if len(roleIds) > 0 {
		// 如果是超级管理员，则返回所有部门
		if isAdmin {
			result, _ = dzhCore.DBM(s.Model).Fields("id").All()
			for _, v := range result {
				vmap := v.Map()
				if vmap["id"] != nil {
					res = append(res, gconv.Uint(vmap["id"]))
				}
			}
		} else {
			// 如果不是超级管理员，则返回角色所在部门
			result, _ = dzhCore.DBM(BaseSysRoleDepartment).Where("roleId IN (?)", roleIds).Fields("departmentId").All()
			for _, v := range result {
				vmap := v.Map()
				if vmap["departmentId"] != nil {
					res = append(res, gconv.Uint(vmap["departmentId"]))
				}
			}
		}

	}
	return
}

// Order 排序部门
func (s *BaseSysDepartmentService) Order(ctx g.Ctx) (err error) {
	r := g.RequestFromCtx(ctx).GetMap()

	type item struct {
		Id       uint64  `json:"id"`
		ParentId *uint64 `json:"parentId,omitempty"`
		OrderNum int32   `json:"orderNum"`
	}

	var data *item

	for _, v := range r {
		err = gconv.Struct(v, &data)
		if err != nil {
			continue
		}
		dzhCore.DBM(s.Model).Where("id = ?", data.Id).Data(data).Update()
	}

	return
}

// NewBaseSysDepartmentService 创建一个BaseSysDepartmentService实例
func NewBaseSysDepartmentService() *BaseSysDepartmentService {
	return &BaseSysDepartmentService{
		Service: &dzhCore.Service{
			Model:       model.NewBaseSysDepartment(),
			ListQueryOp: &dzhCore.QueryOp{},
		},
	}
}
