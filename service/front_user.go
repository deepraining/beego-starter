package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
    "github.com/jinzhu/copier"
)

// 前端用户列表
func FrontUserList(searchKey string, pageSize int64, pageNum int64) (error, *[]models.FrontUser, int64) {
    limit := pageSize
    offset := pageSize * (pageNum - 1)

    query := utils.GetDB().Model(&models.FrontUser{})
    if searchKey != "" {
        query.Where("username like ?", "%"+searchKey+"%")
    }

    list := &[]models.FrontUser{}
    var total int64 = 0
    query.Count(&total)
    result := query.Limit(int(limit)).Offset(int(offset)).Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil, 0
    }
    return nil, list, total
}

// 创建前端用户
func CreateFrontUser(frontUserCreateParam *models.FrontUserCreateParam) (error, int64) {
    frontUser := &models.FrontUser{}
    copier.Copy(frontUser, frontUserCreateParam)
    frontUser.Id = utils.GetUuid().Generate().Int64()
    result := utils.GetDB().Create(frontUser)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 更新前端用户
func UpdateFrontUser(id int64, frontUserCreateParam *models.FrontUserCreateParam) (error, int64) {
    frontUser := &models.FrontUser{}
    copier.Copy(frontUser, frontUserCreateParam)
    frontUser.Id = id
    result := utils.GetDB().Updates(frontUser)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 删除前端用户
func DeleteFrontUser(id int64) (error, int64) {
    result := utils.GetDB().Delete(&models.FrontUser{}, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

