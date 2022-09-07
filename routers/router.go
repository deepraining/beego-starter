package routers

import (
    "github.com/beego/beego/v2/server/web"
    "github.com/deepraining/beego-starter/controllers"
)

func init() {
    web.Router("/", &controllers.AdminController{}, "get:AdminIndex")
    web.Router("/admin/register", &controllers.AdminController{}, "post:AdminRegister")
    web.Router("/admin/login", &controllers.AdminController{}, "post:AdminLogin")
    web.Router("/admin/refreshToken", &controllers.AdminController{}, "get:AdminRefreshToken")
    web.Router("/admin/info", &controllers.AdminController{}, "get:AdminInfo")
    web.Router("/admin/logout", &controllers.AdminController{}, "post:AdminLogout")
    web.Router("/admin/list", &controllers.AdminController{}, "get:AdminUserList")
    web.Router("/admin/update/:id", &controllers.AdminController{}, "post:AdminUserUpdate")
    web.Router("/admin/updatePassword", &controllers.AdminController{}, "post:AdminUserUpdatePassword")
    web.Router("/admin/delete/:id", &controllers.AdminController{}, "post:AdminUserDelete")
    web.Router("/admin/updateStatus/:id", &controllers.AdminController{}, "post:AdminUserUpdateStatus")
    web.Router("/admin/role/update", &controllers.AdminController{}, "post:AdminUserUpdateRole")
    web.Router("/admin/role/:userId", &controllers.AdminController{}, "get:AdminUserRoleList")
    web.Router("/admin/:id", &controllers.AdminController{}, "get:AdminUserItem")

    web.Router("/adminMenu/create", &controllers.AdminMenuController{}, "post:CreateAdminMenu")
    web.Router("/adminMenu/update/:id", &controllers.AdminMenuController{}, "post:UpdateAdminMenu")
    web.Router("/adminMenu/delete/:id", &controllers.AdminMenuController{}, "post:DeleteAdminMenu")
    web.Router("/adminMenu/list/:parentId", &controllers.AdminMenuController{}, "get:AdminMenuList")
    web.Router("/adminMenu/treeList", &controllers.AdminMenuController{}, "get:AdminMenuTreeList")
    web.Router("/adminMenu/updateHidden/:id", &controllers.AdminMenuController{}, "post:AdminMenuUpdateHidden")
    web.Router("/adminMenu/:id", &controllers.AdminMenuController{}, "get:GetAdminMenuItem")

    web.Router("/adminResourceCategory/create", &controllers.AdminResourceCategoryController{}, "post:CreateAdminResourceCategory")
    web.Router("/adminResourceCategory/update/:id", &controllers.AdminResourceCategoryController{}, "post:UpdateAdminResourceCategory")
    web.Router("/adminResourceCategory/delete/:id", &controllers.AdminResourceCategoryController{}, "post:DeleteAdminResourceCategory")
    web.Router("/adminResourceCategory/listAll", &controllers.AdminResourceCategoryController{}, "get:AdminResourceCategoryListAll")

    web.Router("/adminResource/create", &controllers.AdminResourceController{}, "post:CreateAdminResource")
    web.Router("/adminResource/update/:id", &controllers.AdminResourceController{}, "post:UpdateAdminResource")
    web.Router("/adminResource/delete/:id", &controllers.AdminResourceController{}, "post:DeleteAdminResource")
    web.Router("/adminResource/list", &controllers.AdminResourceController{}, "get:AdminResourceList")
    web.Router("/adminResource/listAll", &controllers.AdminResourceController{}, "get:AdminResourceListAll")
    web.Router("/adminResource/:id", &controllers.AdminResourceController{}, "get:GetAdminResourceItem")

    web.Router("/adminRole/create", &controllers.AdminRoleController{}, "post:CreateAdminRole")
    web.Router("/adminRole/update/:id", &controllers.AdminRoleController{}, "post:UpdateAdminRole")
    web.Router("/adminRole/delete", &controllers.AdminRoleController{}, "post:DeleteAdminRole")
    web.Router("/adminRole/list", &controllers.AdminRoleController{}, "get:AdminRoleList")
    web.Router("/adminRole/listAll", &controllers.AdminRoleController{}, "get:AdminRoleListAll")
    web.Router("/adminRole/updateStatus/:roleId", &controllers.AdminRoleController{}, "post:UpdateAdminRoleStatus")
    web.Router("/adminRole/listMenu/:roleId", &controllers.AdminRoleController{}, "get:AdminRoleMenuList")
    web.Router("/adminRole/listResource/:roleId", &controllers.AdminRoleController{}, "get:AdminRoleResourceList")
    web.Router("/adminRole/allocMenu", &controllers.AdminRoleController{}, "post:AllocAdminRoleMenu")
    web.Router("/adminRole/allocResource", &controllers.AdminRoleController{}, "post:AllocAdminRoleResource")

    web.Router("/article/create", &controllers.ArticleController{}, "post:CreateArticle")
    web.Router("/article/update/:id", &controllers.ArticleController{}, "post:UpdateArticle")
    web.Router("/article/delete/:id", &controllers.ArticleController{}, "post:DeleteArticle")
    web.Router("/article/list", &controllers.ArticleController{}, "get:ArticleList")

    web.Router("/frontUser/create", &controllers.FrontUserController{}, "post:CreateFrontUser")
    web.Router("/frontUser/update/:id", &controllers.FrontUserController{}, "post:UpdateFrontUser")
    web.Router("/frontUser/delete/:id", &controllers.FrontUserController{}, "post:DeleteFrontUser")
    web.Router("/frontUser/list", &controllers.FrontUserController{}, "get:FrontUserList")

}
