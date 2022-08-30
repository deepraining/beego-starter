package controllers

import (
    "encoding/json"
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/service"
    "github.com/senntyou/beego-starter/utils"
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
        c.ApiFail("数据解析失败")
    }

    count := service.CreateAdminRole(adminRole)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 修改角色
func (c *AdminRoleController) UpdateAdminRole()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminRole := &models.AdminRole{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminRole)
    if err != nil {
        c.ApiFail("数据解析失败")
    }

    count := service.UpdateAdminRole(id, adminRole)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 根据id批量删除角色
func (c *AdminRoleController) DeleteAdminRole()  {
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("ids")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    count := service.DeleteAdminRole(ids)
    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 获取所有角色列表
func (c *AdminRoleController) AdminRoleListAll()  {
    c.JsonResult(models.SuccessResult(service.AdminRoleListAll()))
}

// 根据角色名称分页获取角色列表
func (c *AdminRoleController) AdminRoleList()  {
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 5)
    keyword := c.Ctx.Input.Query("keyword")

    list, total := service.AdminRoleList(keyword, pageSize, pageNum)
    c.JsonResult(&map[string]interface{}{
        "pageNum": pageNum,
        "pageSize": pageSize,
        "pages": math.Ceil(float64(total)/float64(pageSize)),
        "total": total,
        "list": list,
    })
}

// 获取相应角色权限
func (c *AdminController) AdminRolePermissionList() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    c.JsonResult(models.SuccessResult(service.AdminRolePermissionList(roleId)))
}

// 修改角色权限
func (c *AdminController) UpdateAdminRolePermission() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("permissionIds")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    count := service.UpdateAdminRolePermission(roleId, ids)
    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 修改角色状态
func (c *AdminController) UpdateAdminRoleStatus() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    status := utils.StringToInt64(c.Ctx.Input.Query("status"), 0)
    adminRole := &models.AdminRole{Status: int32(status)}
    count := service.UpdateAdminRole(roleId, adminRole)
    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 获取角色相关菜单
func (c *AdminController) AdminRoleMenuList() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    c.JsonResult(models.SuccessResult(service.AdminMenuListByRoleId(roleId)))
}

// 获取角色相关资源
func (c *AdminController) AdminRoleResourceList() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    c.JsonResult(models.SuccessResult(service.AdminResourceListByRoleId(roleId)))
}

// 给角色分配菜单
func (c *AdminController) AllocAdminRoleMenu() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("menuIds")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    count := service.AllocAdminRoleMenu(roleId, ids)
    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 给角色分配资源
func (c *AdminController) AllocAdminRoleResource() {
    roleId := utils.StringToInt64(c.Ctx.Input.Params()["roleId"], 0)
    // [1,2,3,4]
    idsStr := c.Ctx.Input.Query("resourceIds")
    ids := &[]int64{}
    json.Unmarshal([]byte(idsStr), ids)
    count := service.AllocAdminRoleResource(roleId, ids)
    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}
