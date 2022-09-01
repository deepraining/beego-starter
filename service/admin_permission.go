package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/jinzhu/copier"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
)

// 创建权限
func CreateAdminPermission(adminPermission *models.AdminPermission) (error, int64) {
    adminPermission.Status = 1
    adminPermission.Sort = 0
    result := utils.GetDB().Create(adminPermission)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 更新权限
func UpdateAdminPermission(id int64, adminPermission *models.AdminPermission) (error, int64) {
    adminPermission.Id = id
    result := utils.GetDB().Updates(adminPermission)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 删除权限
func DeleteAdminPermission(ids *[]int64) (error, int64) {
    adminPermission := &models.AdminPermission{}
    result := utils.GetDB().Delete(adminPermission, ids)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}


// 权限列表
func AdminPermissionList() (error, *[]models.AdminPermission) {
    list := &[]models.AdminPermission{}
    result := utils.GetDB().Order("sort desc").Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

// 树形权限列表
func AdminPermissionTreeList() (error, *[]models.AdminPermissionNode)  {
    list := &[]models.AdminPermission{}
    result := utils.GetDB().Order("sort desc").Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }

    treeList := []models.AdminPermissionNode{}
    for _, item := range *list {
        // 根权限
        if item.ParentId == 0 {
            treeList = append(treeList, *convertAdminPermissionNode(&item, list))
        }
    }
    return nil, &treeList
}

func convertAdminPermissionNode(adminPermission *models.AdminPermission, list *[]models.AdminPermission) *models.AdminPermissionNode {
    adminPermissionNode := &models.AdminPermissionNode{}
    copier.Copy(adminPermissionNode, adminPermission)

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
