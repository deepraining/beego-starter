package controllers

import (
    "encoding/json"
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/service"
    "github.com/senntyou/beego-starter/utils"
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
        c.ApiFail("数据解析失败")
    }

    count := service.CreateAdminResource(adminResource)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 修改后台资源
func (c *AdminResourceController) UpdateAdminResource()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminResource := &models.AdminResource{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminResource)
    if err != nil {
        c.ApiFail("数据解析失败")
    }

    count := service.UpdateAdminResource(id, adminResource)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 根据ID获取资源详情
func (c *AdminResourceController) DeleteAdminResource()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    count := service.DeleteAdminResource(id)
    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 分页模糊查询后台资源
func (c *AdminResourceController) AdminResourceList()  {
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 5)
    categoryId := utils.StringToInt64(c.Ctx.Input.Query("categoryId"), 0)
    nameKeyword := c.Ctx.Input.Query("nameKeyword")
    urlKeyword := c.Ctx.Input.Query("urlKeyword")

    list, total := service.AdminResourceList(categoryId, nameKeyword, urlKeyword, pageSize, pageNum)
    c.JsonResult(&map[string]interface{}{
        "pageNum": pageNum,
        "pageSize": pageSize,
        "pages": math.Ceil(float64(total)/float64(pageSize)),
        "total": total,
        "list": list,
    })
}

// 查询所有后台资源
func (c *AdminResourceController) AdminResourceListAll()  {
    c.JsonResult(models.SuccessResult(service.AdminResourceListAll()))
}
