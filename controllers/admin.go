package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
    "math"
)

type AdminController struct {
    BaseController
}

// 主页
func (c *BaseController) AdminIndex()  {
    c.TplName = "index.html"
}

// 用户注册
func (c *AdminController) AdminRegister() {
    adminUserParam := &models.AdminUserParam{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminUserParam)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, adminUser := service.AdminRegister(adminUserParam)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if adminUser == nil {
        c.ApiFail("注册失败")
    }
    c.ApiSucceed(adminUser)
}

// 登录以后返回token
func (c *AdminController) AdminLogin() {
    adminLoginParam := &models.AdminLoginParam{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminLoginParam)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, token := service.AdminLogin(adminLoginParam.Username, adminLoginParam.Password)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if token == "" {
        c.ApiFail("登录失败")
    }
    c.ApiSucceed(&map[string]string{
        "token": token,
        "tokenHead": utils.TokenHead,
    })
}

// 刷新token
func (c *AdminController) AdminRefreshToken() {
    token := c.Ctx.Request.Header.Get(utils.TokenHeaderKey)
    if token == "" {
        c.ApiFail("请先登录")
    }
    err, newToken := service.AdminRefreshToken(token)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if newToken == "" {
        c.ApiFail("token已经过期！")
    }

    c.ApiSucceed(&map[string]string{
        "token": token,
        "tokenHead": utils.TokenHead,
    })
}

// 获取当前登录用户信息
func (c *AdminController) AdminInfo() {
    username := c.Username
    if username == "" {
        c.JsonResult(models.UnauthorizedResult(nil))
    }
    err, adminUser := service.GetAdminUserByUsername(username)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    err2, menus := service.AdminMenuListByUserId(adminUser.Id)
    if err2 != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(&map[string]interface{}{
        "username": username,
        "roles": &[]string{"NONE"},
        "menus": menus,
        "avatar": adminUser.Avatar,
    })
}

// 登出功能
func (c *AdminController) AdminLogout() {
    c.ApiSucceed(nil)
}

// 根据用户名或姓名分页获取用户列表
func (c *AdminController) AdminUserList() {
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 5)
    keyword := c.Ctx.Input.Query("keyword")

    err, list, total := service.AdminUserList(keyword, pageSize, pageNum)
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

// 获取指定用户信息
func (c *AdminController) AdminUserItem() {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    err, data := service.GetAdminUserItem(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 修改指定用户信息
func (c *AdminController) AdminUserUpdate() {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminUser := &models.AdminUser{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminUser)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, count := service.UpdateAdminUser(id, adminUser)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(count)
}

// 修改指定用户密码
func (c *AdminController) AdminUserUpdatePassword() {
    param := &models.UpdateAdminUserPasswordParam{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, param)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err, status := service.UpdateAdminPassword(param)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if status > 0 {
        c.ApiSucceed(status)
    } else if status == -1 {
        c.ApiFail("提交参数不合法")
    } else if status == -2 {
        c.ApiFail("找不到该用户")
    } else if status == -3 {
        c.ApiFail("旧密码错误")
    } else {
        c.ApiFail("")
    }
}

// 删除指定用户信息
func (c *AdminController) AdminUserDelete() {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    err, count := service.DeleteAdminUser(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改帐号状态
func (c *AdminController) AdminUserUpdateStatus() {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    status := utils.StringToInt64(c.Ctx.Input.Query("status"), 0)
    adminUser := &models.AdminUser{
        Status: int32(status),
    }
    err, count := service.UpdateAdminUser(id, adminUser)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 给用户分配角色
func (c *AdminController) AdminUserUpdateRole() {
    userId := utils.StringToInt64(c.Ctx.Input.Query("userId"), 0)
    // [1,2,3,4]
    roleIdsStr := c.Ctx.Input.Query("roleIds")
    roleIds := &[]int64{}
    json.Unmarshal([]byte(roleIdsStr), roleIds)
    err, count := service.UpdateAdminUserRole(userId, roleIds)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 获取指定用户的角色
func (c *AdminController) AdminUserRoleList() {
    userId := utils.StringToInt64(c.Ctx.Input.Params()["userId"], 0)
    err, data := service.AdminRoleListByUserId(userId)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 给用户分配+-权限
func (c *AdminController) AdminUserUpdatePermission() {
    userId := utils.StringToInt64(c.Ctx.Input.Query("userId"), 0)
    // [1,2,3,4]
    permissionIdsStr := c.Ctx.Input.Query("permissionIds")
    permissionIds := &[]int64{}
    json.Unmarshal([]byte(permissionIdsStr), permissionIds)
    err, count := service.UpdateAdminUserPermission(userId, permissionIds)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 获取用户所有权限（包括+-权限）
func (c *AdminController) AdminUserPermissionList() {
    userId := utils.StringToInt64(c.Ctx.Input.Params()["userId"], 0)
    err, data := service.GetAdminPermissionList(userId)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}
