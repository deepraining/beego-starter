package service

import (
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/utils"
)

// 创建权限
func CreateAdminPermission(adminPermission *models.AdminPermission) int64 {
    adminPermission.Status = 1
    adminPermission.Sort = 0
    result := utils.GetDB().Create(adminPermission)
    return result.RowsAffected
}

// 更新权限
func UpdateAdminPermission(id int64, adminPermission *models.AdminPermission) int64 {
    adminPermission.Id = id
    result := utils.GetDB().Updates(adminPermission)
    return result.RowsAffected
}

// 删除权限
func DeleteAdminPermission(ids *[]int64) int64 {
    adminPermission := &models.AdminPermission{}
    result := utils.GetDB().Delete(adminPermission, ids)
    return result.RowsAffected
}


// 权限列表
func AdminPermissionList() *[]models.AdminPermission {
    list := &[]models.AdminPermission{}
    utils.GetDB().Order("sort desc").Find(list)
    return list
}

// 树形权限列表
func AdminPermissionTreeList() *[]models.AdminPermissionNode  {
    list := &[]models.AdminPermission{}
    utils.GetDB().Order("sort desc").Find(list)

    result := []models.AdminPermissionNode{}
    for _, item := range *list {
        // 根权限
        if item.ParentId == 0 {
            result = append(result, *convertAdminPermissionNode(&item, list))
        }
    }
    return &result
}

func convertAdminPermissionNode(adminPermission *models.AdminPermission, list *[]models.AdminPermission) *models.AdminPermissionNode {
    adminPermissionNode := &models.AdminPermissionNode{}
    utils.CopyStructFields(adminPermission, adminPermissionNode)

    children := []models.AdminPermissionNode{}
    for _, item := range *list{
        // 根权限
        if item.ParentId == adminPermissionNode.ParentId {
            children = append(children, *convertAdminPermissionNode(&item, list))
        }
    }
    adminPermissionNode.Children = &children

    return adminPermissionNode
}
