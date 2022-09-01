package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
    "math"
)

type AdminMenuController struct {
    BaseController
}

// 添加后台菜单
func (c *AdminMenuController) CreateAdminMenu()  {
    adminMenu := &models.AdminMenu{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminMenu)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err2, count := service.CreateAdminMenu(adminMenu)
    if err2 != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 修改后台菜单
func (c *AdminMenuController) UpdateAdminMenu()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminMenu := &models.AdminMenu{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminMenu)
    if err != nil {
        logs.Error(err)
        c.ApiFail("数据解析失败")
    }

    err2, count := service.UpdateAdminMenu(id, adminMenu)
    if err2 != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 根据ID获取菜单详情
func (c *AdminMenuController) GetAdminMenuItem()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    err, data := service.GetAdminMenu(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 根据ID删除后台菜单
func (c *AdminMenuController) DeleteAdminMenu()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    err, count := service.DeleteAdminMenu(id)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}

// 分页查询后台菜单
func (c *AdminMenuController) AdminMenuList()  {
    parentId := utils.StringToInt64(c.Ctx.Input.Params()["parentId"], 0)
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 5)

    err, list, total := service.AdminMenuList(parentId, pageSize, pageNum)
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

// 树形结构返回所有菜单列表
func (c *AdminMenuController) AdminMenuTreeList()  {
    err, data := service.AdminMenuTreeList()
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.ApiSucceed(data)
}

// 修改菜单显示状态
func (c *AdminMenuController) AdminMenuUpdateHidden()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    hidden := utils.StringToInt64(c.Ctx.Input.Query("hidden"), 0)

    err, count := service.UpdateAdminMenuHidden(id, hidden)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }

    if count > 0 {
        c.ApiSucceed(count)
    } else {
        c.ApiFail("")
    }
}
