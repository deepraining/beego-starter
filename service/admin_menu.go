package service

import (
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/utils"
)

// 修改菜单层级
func updateAdminMenuLevel(adminMenu *models.AdminMenu) {
    if adminMenu.ParentId == 0 {
        // 没有父菜单时为一级菜单
        adminMenu.Level = 0
    } else {
        parentAdminMenu := &models.AdminMenu{}
        // 有父菜单时选择根据父菜单level设置
        utils.GetDB().First(parentAdminMenu, adminMenu.ParentId)
        if parentAdminMenu.Id > 0 {
            adminMenu.Level = parentAdminMenu.Level + 1
        }else {
            adminMenu.Level = 0
        }
    }
}

// 创建菜单
func CreateAdminMenu(adminMenu *models.AdminMenu) int64 {
    updateAdminMenuLevel(adminMenu)
    result := utils.GetDB().Create(adminMenu)
    return result.RowsAffected
}

// 更新菜单
func UpdateAdminMenu(id int64, adminMenu *models.AdminMenu) int64 {
    adminMenu.Id = id
    updateAdminMenuLevel(adminMenu)
    result := utils.GetDB().Updates(adminMenu)
    return result.RowsAffected
}

// 更新菜单隐藏选项
func UpdateAdminMenuHidden(id int64, hidden int64) int64 {
    adminMenu := &models.AdminMenu{
        Id:     id,
        Hidden: int32(hidden),
    }
    result := utils.GetDB().Updates(adminMenu)
    return result.RowsAffected
}

// 获取菜单
func GetAdminMenu(id int64) *models.AdminMenu {
    adminMenu := &models.AdminMenu{}
    utils.GetDB().First(adminMenu, id)
    if adminMenu.Id > 0 {
        return adminMenu
    }
    return nil
}

// 删除菜单
func DeleteAdminMenu(id int64) int64 {
    adminMenu := &models.AdminMenu{}
    result := utils.GetDB().Delete(adminMenu, id)
    return result.RowsAffected
}


// 菜单列表
func AdminMenuList(parentId int64, pageSize int64, pageNum int64) (*[]models.AdminMenu, int64) {
    limit := pageSize
    offset := pageSize * (pageNum - 1)

    query := utils.GetDB().Where("parentId = ?", parentId)
    list := &[]models.AdminMenu{}
    var total int64 = 0
    query.Count(&total)
    query.Order("sort desc").Limit(int(limit)).Offset(int(offset)).Find(list)
    return list, total
}

// 树形菜单列表
func AdminMenuTreeList() *[]models.AdminMenuNode  {
    list := &[]models.AdminMenu{}
    utils.GetDB().Order("sort desc").Find(list)

    result := []models.AdminMenuNode{}
    for _, item := range *list{
        // 根菜单
        if item.ParentId == 0 {
            result = append(result, *convertAdminMenuNode(&item, list))
        }
    }
    return &result
}

func convertAdminMenuNode(adminMenu *models.AdminMenu, list *[]models.AdminMenu) *models.AdminMenuNode {
    adminMenuNode := &models.AdminMenuNode{}
    utils.CopyStructFields(adminMenu, adminMenuNode)

    children := []models.AdminMenuNode{}
    for _, item := range *list{
        // 根菜单
        if item.ParentId == adminMenuNode.ParentId {
            children = append(children, *convertAdminMenuNode(&item, list))
        }
    }
    adminMenuNode.Children = &children

    return adminMenuNode
}
