package service

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhgo/dzhCore"
	"github.com/gzdzh/dzhgo/modules/base/model"
)

var baseSysUserRole = model.NewBaseSysUserRole()

type BaseSysRoleService struct {
	*dzhCore.Service
}

// ModifyAfter modify after
func (s *BaseSysRoleService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if param["id"] != nil {
		err = s.updatePerms(ctx, gconv.Uint(param["id"]), gconv.SliceUint(param["menuIdList"]), gconv.SliceUint(param["departmentIdList"]))
	}
	return
}

// updatePerms(roleId, menuIdList?, departmentIds = [])
func (s *BaseSysRoleService) updatePerms(ctx context.Context, roleId uint, menuIdList, departmentIds []uint) (err error) {
	// 更新菜单权限
	dzhCore.DBM(model.NewBaseSysRoleMenu()).Where("roleId = ?", roleId).Delete()
	if len(menuIdList) > 0 {
		roleMenuList := make([]g.MapStrAny, len(menuIdList))
		for i, menuId := range menuIdList {
			roleMenuList[i] = g.MapStrAny{
				"roleId": roleId,
				"menuId": menuId,
			}
		}
		dzhCore.DBM(model.NewBaseSysRoleMenu()).Data(roleMenuList).Insert()
	}
	// 更新部门权限
	dzhCore.DBM(model.NewBaseSysRoleDepartment()).Where("roleId = ?", roleId).Delete()
	if len(departmentIds) > 0 {
		roleDepartmentList := make([]g.MapStrAny, len(departmentIds))
		for i, departmentId := range departmentIds {
			roleDepartmentList[i] = g.MapStrAny{
				"roleId":       roleId,
				"departmentId": departmentId,
			}
		}
		dzhCore.DBM(model.NewBaseSysRoleDepartment()).Data(roleDepartmentList).Insert()
	}
	// 刷新权限
	userRoles, err := dzhCore.DBM(model.NewBaseSysUserRole()).Where("roleId = ?", roleId).All()
	if err != nil {
		return
	}
	baseSysPermsService := NewBaseSysPermsService()
	for _, v := range userRoles {
		vmap := v.Map()
		if vmap["userId"] != nil {
			baseSysPermsService.RefreshPerms(ctx, gconv.Uint(vmap["userId"]))
		}
	}

	return

}

// GetByUser get array  roleId by userId
func (s *BaseSysRoleService) GetByUser(userId uint) []string {
	var (
		roles []string
	)
	res, _ := dzhCore.DBM(baseSysUserRole).Where("userId = ?", userId).Array("roleId")
	for _, v := range res {
		roles = append(roles, gconv.String(v))
	}
	return roles
}

// BaseSysRoleService Info 方法重构
func (s *BaseSysRoleService) ServiceInfo(ctx context.Context, req *dzhCore.InfoReq) (data interface{}, err error) {
	info, err := dzhCore.DBM(s.Model).Where("id = ?", req.Id).One()
	if err != nil {
		return nil, err
	}
	if !info.IsEmpty() {
		var menus gdb.Result
		if req.Id == 1 {
			menus, err = dzhCore.DBM(model.NewBaseSysMenu()).All()
			if err != nil {
				return nil, err
			}
		} else {
			menus, err = dzhCore.DBM(model.NewBaseSysRoleMenu()).Where("roleId = ?", req.Id).All()
			if err != nil {
				return nil, err
			}
		}
		menuIdList := garray.NewIntArray()
		for _, v := range menus {
			menuIdList.Append(gconv.Int(v["menuId"]))
		}
		var departments gdb.Result
		if req.Id == 1 {
			departments, err = dzhCore.DBM(model.NewBaseSysRoleDepartment()).All()
			if err != nil {
				return nil, err
			}
		} else {
			departments, err = dzhCore.DBM(model.NewBaseSysRoleDepartment()).Where("roleId = ?", req.Id).All()
			if err != nil {
				return nil, err
			}
		}
		departmentIdList := garray.NewIntArray()
		for _, v := range departments {
			departmentIdList.Append(gconv.Int(v["departmentId"]))
		}
		result := gconv.Map(info)
		result["menuIdList"] = menuIdList.Slice()
		result["departmentIdList"] = departmentIdList.Slice()
		data = result
		return
	}
	data = g.Map{}
	return
}

// NewBaseSysRoleService create a new BaseSysRoleService
func NewBaseSysRoleService() *BaseSysRoleService {
	return &BaseSysRoleService{
		Service: &dzhCore.Service{
			Model: model.NewBaseSysRole(),
			ListQueryOp: &dzhCore.QueryOp{
				Where: func(ctx context.Context) [][]interface{} {
					var (
						admin   = dzhCore.GetAdmin(ctx)
						userId  = admin.UserId
						roleIds = garray.NewIntArrayFromCopy(gconv.Ints(admin.RoleIds))
					)
					return [][]interface{}{
						{"label != ?", g.Slice{"admin"}, true},
						{"(userId=? or id in (?))", g.Slice{userId, admin.RoleIds}, !roleIds.Contains(1)},
					}
				},
			},
			PageQueryOp: &dzhCore.QueryOp{
				KeyWordField: []string{"name", "label"},
				AddOrderby:   map[string]string{},
				Where: func(ctx context.Context) [][]interface{} {
					var (
						admin   = dzhCore.GetAdmin(ctx)
						userId  = admin.UserId
						roleIds = garray.NewIntArrayFromCopy(gconv.Ints(admin.RoleIds))
					)
					return [][]interface{}{
						{"label != ?", g.Slice{"admin"}, true},
						{"(userId=? or id in (?))", g.Slice{userId, admin.RoleIds}, !roleIds.Contains(1)},
					}
				},
			},
			InsertParam: func(ctx context.Context) map[string]interface{} {
				return g.Map{"userId": dzhCore.GetAdmin(ctx).UserId}
			},
			UniqueKey: map[string]string{
				"name":  "角色名称不能重复",
				"label": "角色标识不能重复",
			},
		},
	}
}
