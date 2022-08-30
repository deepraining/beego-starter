package service

import (
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/utils"
)

// 创建资源
func CreateAdminResource(adminResource *models.AdminResource) int64 {
    result := utils.GetDB().Create(adminResource)
    return result.RowsAffected
}

// 更新资源
func UpdateAdminResource(id int64, adminResource *models.AdminResource) int64 {
    adminResource.Id = id
    result := utils.GetDB().Updates(adminResource)
    DelAdminResourceListByResourceCache(id);
    return result.RowsAffected
}

// 获取资源
func GetAdminResource(id int64) *models.AdminResource {
    adminResource := &models.AdminResource{}
    utils.GetDB().First(adminResource, id)
    if adminResource.Id > 0 {
        return adminResource
    }
    return nil
}

// 删除资源
func DeleteAdminResource(id int64) int64 {
    adminResource := &models.AdminResource{}
    result := utils.GetDB().Delete(adminResource, id)
    DelAdminResourceListByResourceCache(id);
    return result.RowsAffected
}

// 资源列表
func AdminResourceList(categoryId int64, nameKeyword string, urlKeyword string, pageSize int64, pageNum int64) (*[]models.AdminResource, int64) {
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
    query.Limit(int(limit)).Offset(int(offset)).Find(list)
    return list, total
}

// 资源列表
func AdminResourceListAll() *[]models.AdminResource {
    list := &[]models.AdminResource{}
    utils.GetDB().Find(list)
    return list
}
