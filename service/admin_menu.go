package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
    "github.com/jinzhu/copier"
)

// 修改菜单层级
func updateAdminMenuLevel(adminMenu *models.AdminMenu) error {
    if adminMenu.ParentId == 0 {
        // 没有父菜单时为一级菜单
        adminMenu.Level = 0
    } else {
        parentAdminMenu := &models.AdminMenu{}
        // 有父菜单时选择根据父菜单level设置
        result := utils.GetDB().Find(parentAdminMenu, adminMenu.ParentId)
        if result.Error != nil {
            logs.Error(result.Error)
            return result.Error
        }
        if parentAdminMenu.Id > 0 {
            adminMenu.Level = parentAdminMenu.Level + 1
        }else {
            adminMenu.Level = 0
        }
    }
    return nil
}

// 创建菜单
func CreateAdminMenu(adminMenu *models.AdminMenu) (error, int64) {
    err := updateAdminMenuLevel(adminMenu)
    if err != nil {
        logs.Error(err)
        return err, 0
    }
    result := utils.GetDB().Create(adminMenu)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 更新菜单
func UpdateAdminMenu(id int64, adminMenu *models.AdminMenu) (error, int64) {
    adminMenu.Id = id
    err := updateAdminMenuLevel(adminMenu)
    if err != nil {
        logs.Error(err)
        return err, 0
    }
    result := utils.GetDB().Updates(adminMenu)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 更新菜单隐藏选项
func UpdateAdminMenuHidden(id int64, hidden int64) (error, int64) {
    adminMenu := &models.AdminMenu{
        Id:     id,
        Hidden: int32(hidden),
    }
    result := utils.GetDB().Updates(adminMenu)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 获取菜单
func GetAdminMenu(id int64) (error, *models.AdminMenu) {
    adminMenu := &models.AdminMenu{}
    result := utils.GetDB().Find(adminMenu, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    if adminMenu.Id > 0 {
        return nil, adminMenu
    }
    return nil, nil
}

// 删除菜单
func DeleteAdminMenu(id int64) (error, int64) {
    result := utils.GetDB().Delete(&models.AdminMenu{}, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}


// 菜单列表
func AdminMenuList(parentId int64, pageSize int64, pageNum int64) (error, *[]models.AdminMenu, int64) {
    limit := pageSize
    offset := pageSize * (pageNum - 1)

    query := utils.GetDB().Model(&models.AdminMenu{}).Where("parent_id = ?", parentId)
    list := &[]models.AdminMenu{}
    var total int64 = 0
    query.Count(&total)
    result := query.Order("sort desc").Limit(int(limit)).Offset(int(offset)).Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil, 0
    }
    return nil, list, total
}

// 树形菜单列表
func AdminMenuTreeList() (error, *[]models.AdminMenuNode)  {
    list := &[]models.AdminMenu{}
    result := utils.GetDB().Order("sort desc").Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }

    nodes := []models.AdminMenuNode{}
    for _, item := range *list{
        // 根菜单
        if item.ParentId == 0 {
            nodes = append(nodes, *convertAdminMenuNode(&item, list))
        }
    }
    return nil, &nodes
}

func convertAdminMenuNode(adminMenu *models.AdminMenu, list *[]models.AdminMenu) *models.AdminMenuNode {
    adminMenuNode := &models.AdminMenuNode{}
    copier.Copy(adminMenuNode, adminMenu)

    children := []models.AdminMenuNode{}
    for _, item := range *list{
        // 根菜单
        if item.ParentId == adminMenuNode.Id {
            children = append(children, *convertAdminMenuNode(&item, list))
        }
    }
    adminMenuNode.Children = &children

    return adminMenuNode
}
