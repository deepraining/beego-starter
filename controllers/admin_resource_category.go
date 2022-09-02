package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
)

type AdminResourceCategoryController struct {
    BaseController
}

// 添加后台资源分类
func (c *AdminResourceCategoryController) CreateAdminResourceCategory()  {
    adminResourceCategory := &models.AdminResourceCategory{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminResourceCategory)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.CreateAdminResourceCategory(adminResourceCategory)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改后台资源分类
func (c *AdminResourceCategoryController) UpdateAdminResourceCategory()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    adminResourceCategory := &models.AdminResourceCategory{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminResourceCategory)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.UpdateAdminResourceCategory(id, adminResourceCategory)
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
func (c *AdminResourceCategoryController) DeleteAdminResourceCategory()  {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    err, count := service.DeleteAdminResourceCategory(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 查询所有后台资源分类
func (c *AdminResourceCategoryController) AdminResourceCategoryListAll()  {
    err, data := service.AdminResourceCategoryListAll()
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}
