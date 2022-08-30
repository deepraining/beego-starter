package controllers

import (
    "encoding/json"
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/service"
    "github.com/senntyou/beego-starter/utils"
)

type AdminResourceCategoryController struct {
    BaseController
}

// 添加后台资源分类
func (c *AdminResourceCategoryController) CreateAdminResourceCategory()  {
    adminResourceCategory := &models.AdminResourceCategory{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminResourceCategory)
    if err != nil {
        c.ApiFail("数据解析失败")
    }

    count := service.CreateAdminResourceCategory(adminResourceCategory)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 修改后台资源分类
func (c *AdminResourceCategoryController) UpdateAdminResourceCategory()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminResourceCategory := &models.AdminResourceCategory{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminResourceCategory)
    if err != nil {
        c.ApiFail("数据解析失败")
    }

    count := service.UpdateAdminResourceCategory(id, adminResourceCategory)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 根据ID删除后台资源
func (c *AdminResourceCategoryController) DeleteAdminResourceCategory()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    count := service.DeleteAdminResourceCategory(id)
    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 查询所有后台资源分类
func (c *AdminResourceCategoryController) AdminResourceCategoryListAll()  {
    c.JsonResult(models.SuccessResult(service.AdminResourceCategoryListAll()))
}
