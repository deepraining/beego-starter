package controllers

import (
    "encoding/json"
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/service"
    "github.com/senntyou/beego-starter/utils"
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
        c.ApiFail("数据解析失败")
    }

    count := service.CreateAdminMenu(adminMenu)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 修改后台菜单
func (c *AdminMenuController) UpdateAdminMenu()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    adminMenu := &models.AdminMenu{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, adminMenu)
    if err != nil {
        c.ApiFail("数据解析失败")
    }

    count := service.UpdateAdminMenu(id, adminMenu)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 根据ID获取菜单详情
func (c *AdminMenuController) GetAdminMenuItem()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    c.JsonResult(models.SuccessResult(service.GetAdminMenu(id)))
}

// 根据ID删除后台菜单
func (c *AdminMenuController) DeleteAdminMenu()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    count := service.DeleteAdminMenu(id)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}

// 分页查询后台菜单
func (c *AdminMenuController) AdminMenuList()  {
    parentId := utils.StringToInt64(c.Ctx.Input.Params()["parentId"], 0)
    pageNum := utils.StringToInt64(c.Ctx.Input.Query("pageNum"), 1)
    pageSize := utils.StringToInt64(c.Ctx.Input.Query("pageSize"), 5)

    list, total := service.AdminMenuList(parentId, pageSize, pageNum)
    c.JsonResult(&map[string]interface{}{
        "pageNum": pageNum,
        "pageSize": pageSize,
        "pages": math.Ceil(float64(total)/float64(pageSize)),
        "total": total,
        "list": list,
    })
}

// 树形结构返回所有菜单列表
func (c *AdminMenuController) AdminMenuTreeList()  {
    c.JsonResult(models.SuccessResult(service.AdminMenuTreeList()))
}

// 修改菜单显示状态
func (c *AdminMenuController) AdminMenuUpdateHidden()  {
    id := utils.StringToInt64(c.Ctx.Input.Params()["id"], 0)
    hidden := utils.StringToInt64(c.Ctx.Input.Query("hidden"), 0)

    count := service.UpdateAdminMenuHidden(id, hidden)

    if count > 0 {
        c.JsonResult(models.SuccessResult(count))
    } else {
        c.JsonResult(models.FailedResult())
    }
}
