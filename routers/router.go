package routers

import (
    "github.com/beego/beego/v2/server/web"
    "github.com/senntyou/beego-starter/controllers"
)

func init() {
    web.Router("/", &controllers.BaseController{})
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
    web.Router("/admin/permission/update", &controllers.AdminController{}, "post:AdminUserUpdatePermission")
    web.Router("/admin/permission/:userId", &controllers.AdminController{}, "get:AdminUserPermissionList")
    web.Router("/admin/:id", &controllers.AdminController{}, "get:AdminUserItem")

    
}
