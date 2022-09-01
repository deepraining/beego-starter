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
    Username string
    Password string
    Avatar string
    Email string
    Nickname string
    Note string
}

type AdminLoginParam struct {
    Username string
    Password string
}

type UpdateAdminUserPasswordParam struct {
    Username string
    OldPassword string
    NewPassword string
}
