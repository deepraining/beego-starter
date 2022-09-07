package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
    "math"
)

type ArticleController struct {
    BaseController
}

// 添加文章
func (c *ArticleController) CreateArticle()  {
    articleCreateParam := &models.ArticleCreateParam{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, articleCreateParam)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.CreateArticle(articleCreateParam)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改文章
func (c *ArticleController) UpdateArticle()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    articleCreateParam := &models.ArticleCreateParam{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, articleCreateParam)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.UpdateArticle(id, articleCreateParam)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 根据ID删除后台资源
func (c *ArticleController) DeleteArticle()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    err, count := service.DeleteArticle(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 查询所有文章
func (c *ArticleController) ArticleList()  {
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 10)
    searchKey := c.Ctx.Input.Query("searchKey")
    err, list, total := service.ArticleList(searchKey, pageSize, pageNum)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(&map[string]interface{}{
        "pageNum": pageNum,
        "pageSize": pageSize,
        "pages": math.Ceil(float64(total)/float64(pageSize)),
        "total": total,
        "list": list,
    })
}
