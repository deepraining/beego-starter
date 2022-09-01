package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
)

// 创建角色
func CreateAdminRole(adminRole *models.AdminRole) (error, int64) {
    adminRole.UserCount = 0
    adminRole.Sort = 0
    result := utils.GetDB().Create(adminRole)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 更新角色
func UpdateAdminRole(id int64, adminRole *models.AdminRole) (error, int64) {
    adminRole.Id = id
    result := utils.GetDB().Updates(adminRole)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 删除角色
func DeleteAdminRole(ids *[]int64) (error, int64) {
    adminRole := &models.AdminRole{}
    result := utils.GetDB().Delete(adminRole, ids)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    DelAdminResourceListByRoleIdsCache(ids);
    return nil, result.RowsAffected
}


// 角色的权限列表
func AdminRolePermissionList(roleId int64) (error, *[]models.AdminPermission) {
    list := &[]models.AdminPermission{}
    result:=utils.GetDB().Raw(`
SELECT
  p.*
FROM
  admin_role_permission_relation r
  LEFT JOIN admin_permission p ON r.permission_id = p.id
WHERE
  r.role_id = ?
`, roleId).Scan(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

// 更新角色的权限
func UpdateAdminRolePermission(roleId int64, permissionIds *[]int64) (error, int64) {
    // 先删除原有关系
    result0 := utils.GetDB().Where("roleId = ?", roleId).Delete(&models.AdminRolePermissionRelation{})
    if result0.Error != nil {
        logs.Error(result0.Error)
        return result0.Error, 0
    }

    // 批量插入新关系
    relationList := []models.AdminRolePermissionRelation{}
    for _, item := range *permissionIds {
        relationList = append(relationList, models.AdminRolePermissionRelation{
            RoleId: roleId,
            PermissionId: item,
        })
    }
    result := utils.GetDB().Create(relationList)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 角色所有列表
func AdminRoleListAll() (error, *[]models.AdminRole) {
    list := &[]models.AdminRole{}
    result := utils.GetDB().Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

// 角色列表
func AdminRoleList(keyword string, pageSize int64, pageNum int64) (error, *[]models.AdminRole, int64) {
    limit := pageSize
    offset := pageSize * (pageNum - 1)

    query := utils.GetDB()
    if keyword != "" {
        query.Where("name like ?", "%"+ keyword +"%")
    }
    list := &[]models.AdminRole{}
    var total int64 = 0
    query.Count(&total)
    result := query.Limit(int(limit)).Offset(int(offset)).Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil, 0
    }
    return nil, list, total
}

// 通过角色ID获取菜单列表
func AdminMenuListByRoleId(roleId int64) (error, *[]models.AdminMenu) {
    list := &[]models.AdminMenu{}
    result := utils.GetDB().Raw(`
SELECT
  m.*
FROM
  admin_role_menu_relation rmr
    LEFT JOIN admin_menu m ON rmr.menu_id = m.id
WHERE
  rmr.role_id = ?
    AND m.id IS NOT NULL
GROUP BY
  m.id
`, roleId).Scan(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

// 通过角色ID获取资源列表
func AdminResourceListByRoleId(roleId int64) (error, *[]models.AdminResource) {
    list := &[]models.AdminResource{}
    result := utils.GetDB().Raw(`
SELECT
  r.*
FROM
  admin_role_resource_relation rrr
    LEFT JOIN admin_resource r ON rrr.resource_id = r.id
WHERE
  rrr.role_id = ?
    AND r.id IS NOT NULL
GROUP BY
  r.id
`, roleId).Scan(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

// 分配角色的菜单
func AllocAdminRoleMenu(roleId int64, menuIds *[]int64) (error, int64) {
    // 先删除原有关系
    utils.GetDB().Where("roleId = ?", roleId).Delete(&models.AdminRoleMenuRelation{})

    // 批量插入新关系
    relationList := []models.AdminRoleMenuRelation{}
    for _, item := range *menuIds {
        relationList = append(relationList, models.AdminRoleMenuRelation{
            RoleId: roleId,
            MenuId: item,
        })
    }
    result := utils.GetDB().Create(relationList)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 分配角色的资源
func AllocAdminRoleResource(roleId int64, resourceIds *[]int64) (error, int64) {
    // 先删除原有关系
    utils.GetDB().Where("roleId = ?", roleId).Delete(&models.AdminRoleResourceRelation{})

    // 批量插入新关系
    relationList := []models.AdminRoleResourceRelation{}
    for _, item := range *resourceIds{
        relationList = append(relationList, models.AdminRoleResourceRelation{
            RoleId: roleId,
            ResourceId: item,
        })
    }
    result := utils.GetDB().Create(relationList)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}
