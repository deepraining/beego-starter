package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
)

// 创建资源
func CreateAdminResource(adminResource *models.AdminResource) (error, int64) {
    result := utils.GetDB().Create(adminResource)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 更新资源
func UpdateAdminResource(id int64, adminResource *models.AdminResource) (error, int64) {
    adminResource.Id = id
    result := utils.GetDB().Updates(adminResource)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    DelAdminResourceListByResourceCache(id);
    return nil, result.RowsAffected
}

// 获取资源
func GetAdminResource(id int64) (error, *models.AdminResource) {
    adminResource := &models.AdminResource{}
    result := utils.GetDB().First(adminResource, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    if adminResource.Id > 0 {
        return nil, adminResource
    }
    return nil, nil
}

// 删除资源
func DeleteAdminResource(id int64) (error, int64) {
    adminResource := &models.AdminResource{}
    result := utils.GetDB().Delete(adminResource, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    DelAdminResourceListByResourceCache(id);
    return nil, result.RowsAffected
}

// 资源列表
func AdminResourceList(categoryId int64, nameKeyword string, urlKeyword string, pageSize int64, pageNum int64) (error, *[]models.AdminResource, int64) {
    limit := pageSize
    offset := pageSize * (pageNum - 1)

    query := utils.GetDB()
    if categoryId != 0 {
        query.Where("categoryId = ?", categoryId)
    }
    if nameKeyword != "" {
        query.Where("name like ?", "%"+nameKeyword+"%")
    }
    if urlKeyword != "" {
        query.Where("url like ?", "%"+urlKeyword+"%")
    }

    list := &[]models.AdminResource{}
    var total int64 = 0
    query.Count(&total)
    result := query.Limit(int(limit)).Offset(int(offset)).Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil, 0
    }
    return nil, list, total
}

// 资源列表
func AdminResourceListAll() (error, *[]models.AdminResource) {
    list := &[]models.AdminResource{}
    result := utils.GetDB().Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}
