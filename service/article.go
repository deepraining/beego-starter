package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
    "github.com/jinzhu/copier"
)

// 文章列表
func ArticleList(searchKey string, pageSize int64, pageNum int64) (error, *[]models.Article, int64) {
    limit := pageSize
    offset := pageSize * (pageNum - 1)

    query := utils.GetDB().Model(&models.Article{})
    if searchKey != "" {
        query.Where("title like ?", "%"+searchKey+"%")
    }

    list := &[]models.Article{}
    var total int64 = 0
    query.Count(&total)
    result := query.Limit(int(limit)).Offset(int(offset)).Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil, 0
    }
    return nil, list, total
}

// 创建文章
func CreateArticle(articleCreateParam *models.ArticleCreateParam) (error, int64) {
    article := &models.Article{}
    copier.Copy(article, articleCreateParam)
    article.Id = utils.GetUuid().Generate().Int64()
    article.FrontUserId = -1
    result := utils.GetDB().Create(article)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 更新文章
func UpdateArticle(id int64, articleCreateParam *models.ArticleCreateParam) (error, int64) {
    article := &models.Article{}
    copier.Copy(article, articleCreateParam)
    article.Id = id
    result := utils.GetDB().Updates(article)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

// 删除文章
func DeleteArticle(id int64) (error, int64) {
    result := utils.GetDB().Delete(&models.Article{}, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
}

