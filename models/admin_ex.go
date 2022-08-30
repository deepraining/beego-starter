package models

type AdminMenuNode struct {
    AdminMenu
    Children *[]AdminMenuNode
}

type AdminPermissionNode struct {
    AdminPermission
    Children *[]AdminPermissionNode
}

type AdminUserDetails struct {
    AdminUser
    ResourceList *[]AdminResource
}

type AdminUserParam struct {
    Username string `valid:"Required"`
    Password string `valid:"Required"`
    Avatar string
    Email string `valid:"Email"`
    Nickname string
    Note string
}

type AdminLoginParam struct {
    Username string `valid:"Required"`
    Password string `valid:"Required"`
}

type UpdateAdminUserPasswordParam struct {
    Username string `valid:"Required"`
    OldPassword string `valid:"Required"`
    NewPassword string `valid:"Required"`
}
