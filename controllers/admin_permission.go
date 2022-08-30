package controllers

import (
    "encoding/json"
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/service"
    "github.com/senntyou/beego-starter/utils"
)

type AdminPermissionController struct {
    BaseController
}

// 添加权限
func (c *AdminPermissionController) CreateAdminPermission()  {
    adminPermission := &models.AdminPermission{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminPermission)
    if err != nil {
        c.ApiFail("数据解析失败")
    }

    count := service.CreateAdminPermission(adminPermission)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 修改权限
func (c *AdminPermissionController) UpdateAdminPermission()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminPermission := &models.AdminPermission{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminPermission)
    if err != nil {
        c.ApiFail("数据解析失败")
    }

    count := service.UpdateAdminPermission(id, adminPermission)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 根据id批量删除权限
func (c *AdminPermissionController) DeleteAdminPermission()  {
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("ids")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    count := service.DeleteAdminPermission(ids)
    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 获取所有权限列表
func (c *AdminPermissionController) AdminPermissionList()  {
    c.JsonResult(models.SuccessResult(service.AdminPermissionList()))
}

// 树形结构返回所有权限列表
func (c *AdminPermissionController) AdminPermissionTreeList()  {
    c.JsonResult(models.SuccessResult(service.AdminPermissionTreeList()))
}
