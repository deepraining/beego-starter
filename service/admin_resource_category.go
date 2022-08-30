package service

import (
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/utils"
)

// 资源分类列表
func AdminResourceCategoryListAll() *[]models.AdminResourceCategory {
    list := &[]models.AdminResourceCategory{}
    utils.GetDB().Order("sort desc").Find(list)
    return list
}

// 创建资源分类
func CreateAdminResourceCategory(adminResourceCategory *models.AdminResourceCategory) int64 {
    result := utils.GetDB().Create(adminResourceCategory)
    return result.RowsAffected
}

// 更新资源分类
func UpdateAdminResourceCategory(id int64, adminResourceCategory *models.AdminResourceCategory) int64 {
    adminResourceCategory.Id = id
    result := utils.GetDB().Updates(adminResourceCategory)
    return result.RowsAffected
}

// 删除资源分类
func DeleteAdminResourceCategory(id int64) int64 {
    adminResourceCategory := &models.AdminResourceCategory{}
    result := utils.GetDB().Delete(adminResourceCategory, id)
    return result.RowsAffected
}

