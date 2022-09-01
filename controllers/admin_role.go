package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
    "math"
)

type AdminRoleController struct {
    BaseController
}

// 添加角色
func (c *AdminRoleController) CreateAdminRole()  {
    adminRole := &models.AdminRole{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminRole)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.CreateAdminRole(adminRole)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改角色
func (c *AdminRoleController) UpdateAdminRole()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminRole := &models.AdminRole{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminRole)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.UpdateAdminRole(id, adminRole)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 根据id批量删除角色
func (c *AdminRoleController) DeleteAdminRole()  {
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("ids")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    err, count := service.DeleteAdminRole(ids)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 获取所有角色列表
func (c *AdminRoleController) AdminRoleListAll()  {
    err, data := service.AdminRoleListAll()
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 根据角色名称分页获取角色列表
func (c *AdminRoleController) AdminRoleList()  {
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 5)
    keyword := c.Ctx.Input.Query("keyword")

    err, list, total := service.AdminRoleList(keyword, pageSize, pageNum)
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

// 获取相应角色权限
func (c *AdminRoleController) AdminRolePermissionList() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    err, data := service.AdminRolePermissionList(roleId)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 修改角色权限
func (c *AdminRoleController) UpdateAdminRolePermission() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("permissionIds")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    err, count := service.UpdateAdminRolePermission(roleId, ids)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改角色状态
func (c *AdminRoleController) UpdateAdminRoleStatus() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    status := utils.StringToInt64(c.Ctx.Input.Query("status"), 0)
    adminRole := &models.AdminRole{Status: int32(status)}
    err, count := service.UpdateAdminRole(roleId, adminRole)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 获取角色相关菜单
func (c *AdminRoleController) AdminRoleMenuList() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    err, data := service.AdminMenuListByRoleId(roleId)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 获取角色相关资源
func (c *AdminRoleController) AdminRoleResourceList() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    err, data := service.AdminResourceListByRoleId(roleId)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 给角色分配菜单
func (c *AdminRoleController) AllocAdminRoleMenu() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("menuIds")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    err, count := service.AllocAdminRoleMenu(roleId, ids)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 给角色分配资源
func (c *AdminRoleController) AllocAdminRoleResource() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("resourceIds")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    err, count := service.AllocAdminRoleResource(roleId, ids)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}
