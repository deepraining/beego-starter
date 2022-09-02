package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
)

// 资源分类列表
func AdminResourceCategoryListAll() (error, *[]models.AdminResourceCategory) {
    list := &[]models.AdminResourceCategory{}
    result := utils.GetDB().Order("sort desc").Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

// 创建资源分类
func CreateAdminResourceCategory(adminResourceCategory *models.AdminResourceCategory) (error, int64) {
    result := utils.GetDB().Create(adminResourceCategory)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 更新资源分类
func UpdateAdminResourceCategory(id int64, adminResourceCategory *models.AdminResourceCategory) (error, int64) {
    adminResourceCategory.Id = id
    result := utils.GetDB().Updates(adminResourceCategory)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 删除资源分类
func DeleteAdminResourceCategory(id int64) (error, int64) {
    result := utils.GetDB().Delete(&models.AdminResourceCategory{}, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

