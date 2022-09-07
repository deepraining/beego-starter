package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
    "math"
    "strconv"
    "strings"
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
        "tokenHead": utils.JwtTokenHead,
    })
}

// 刷新token
func (c *AdminController) AdminRefreshToken() {
    token := c.Ctx.Request.Header.Get(utils.JwtTokenHeaderKey)
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
        "tokenHead": utils.JwtTokenHead,
    })
}

// 获取当前登录用户信息
func (c *AdminController) AdminInfo() {
    username := c.Username
    if username == "" {
        c.JsonResult(models.UnauthorizedResult())
    }
    err, adminUser := service.GetAdminUserByUsername(username)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    err, menus := service.AdminMenuListByUserId(adminUser.Id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(&map[string]interface{}{
        "username": adminUser.Username,
        "nickname": adminUser.Nickname,
        "avatar": adminUser.Avatar,
        "roles": &[]string{"NONE"},
        "menus": menus,
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
    searchKey := c.Ctx.Input.Query("searchKey")

    err, list, total := service.AdminUserList(searchKey, pageSize, pageNum)
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
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    err, data := service.GetAdminUserItem(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 修改指定用户信息
func (c *AdminController) AdminUserUpdate() {
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
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
    param := &models.AdminUpdatePasswordParam{}
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
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
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
    id := utils.StringToInt64(c.Ctx.Input.Param(":id"), 0)
    if id == 0 {
        c.ApiFail("参数错误")
    }
    status := utils.StringToInt64(c.Ctx.Input.Query("status"), 0)

    adminUser := &map[string]interface{}{
        "Status": int64(status),
    }
    err, count := service.UpdateAdminUserByMap(id, adminUser)
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
    // 1,2,3,4
    idsStr := c.Ctx.Input.Query("roleIds")
    if idsStr == "" {
        c.ApiFail("参数错误")
    }
    idStrList := strings.Split(idsStr, ",")
    ids := []int64{}
    for _, idStr := range idStrList{
        id, _ := strconv.Atoi(idStr)
        ids = append(ids, int64(id))
    }
    err, count := service.UpdateAdminUserRole(userId, &ids)
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
    userId := utils.StringToInt64(c.Ctx.Input.Param(":userId"), 0)
    if userId == 0 {
        c.ApiFail("参数错误")
    }
    err, data := service.AdminRoleListByUserId(userId)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}
