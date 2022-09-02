package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
    "math"
)

type AdminResourceController struct {
    BaseController
}

// 添加后台资源
func (c *AdminResourceController) CreateAdminResource()  {
    adminResource := &models.AdminResource{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminResource)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.CreateAdminResource(adminResource)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改后台资源
func (c *AdminResourceController) UpdateAdminResource()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    adminResource := &models.AdminResource{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminResource)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.UpdateAdminResource(id, adminResource)
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
func (c *AdminResourceController) DeleteAdminResource()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    err, count := service.DeleteAdminResource(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 根据ID获取资源详情
func (c *AdminResourceController) GetAdminResourceItem()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    err, data := service.GetAdminResource(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 分页模糊查询后台资源
func (c *AdminResourceController) AdminResourceList()  {
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 5)
    categoryId := utils.StringToInt64(c.Ctx.Input.Query("categoryId"), 0)
    nameKeyword := c.Ctx.Input.Query("nameKeyword")
    urlKeyword := c.Ctx.Input.Query("urlKeyword")

    err, list, total := service.AdminResourceList(categoryId, nameKeyword, urlKeyword, pageSize, pageNum)
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

// 查询所有后台资源
func (c *AdminResourceController) AdminResourceListAll()  {
    err, data := service.AdminResourceListAll()
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}
