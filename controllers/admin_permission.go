package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
)

type AdminPermissionController struct {
    BaseController
}

// 添加权限
func (c *AdminPermissionController) CreateAdminPermission()  {
    adminPermission := &models.AdminPermission{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminPermission)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.CreateAdminPermission(adminPermission)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改权限
func (c *AdminPermissionController) UpdateAdminPermission()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminPermission := &models.AdminPermission{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminPermission)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err2, count := service.UpdateAdminPermission(id, adminPermission)
    if err2 != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 根据id批量删除权限
func (c *AdminPermissionController) DeleteAdminPermission()  {
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("ids")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    err, count := service.DeleteAdminPermission(ids)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 获取所有权限列表
func (c *AdminPermissionController) AdminPermissionList()  {
    err, data := service.AdminPermissionList()
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 树形结构返回所有权限列表
func (c *AdminPermissionController) AdminPermissionTreeList()  {
    err, data := service.AdminPermissionTreeList()
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}
