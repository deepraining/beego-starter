package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
    "math"
)

type FrontUserController struct {
    BaseController
}

// 添加前端用户
func (c *FrontUserController) CreateFrontUser()  {
    frontUserCreateParam := &models.FrontUserCreateParam{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, frontUserCreateParam)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.CreateFrontUser(frontUserCreateParam)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改前端用户
func (c *FrontUserController) UpdateFrontUser()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    frontUserCreateParam := &models.FrontUserCreateParam{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, frontUserCreateParam)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.UpdateFrontUser(id, frontUserCreateParam)
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
func (c *FrontUserController) DeleteFrontUser()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    err, count := service.DeleteFrontUser(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 查询所有前端用户
func (c *FrontUserController) FrontUserList()  {
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 10)
    searchKey := c.Ctx.Input.Query("searchKey")
    err, list, total := service.FrontUserList(searchKey, pageSize, pageNum)
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
