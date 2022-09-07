package models

type AdminMenuNode struct {
    AdminMenu
    Children *[]AdminMenuNode `json:"children"`
}

type AdminUserDetails struct {
    AdminUser
    ResourceList *[]AdminResource `json:"resourceList"`
}

type AdminUserParam struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Avatar string `json:"avatar"`
    Email string `json:"email"`
    Nickname string `json:"nickname"`
    Note string `json:"note"`
}

type AdminLoginParam struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type AdminUpdatePasswordParam struct {
    Username string `json:"username"`
    OldPassword string `json:"oldPassword"`
    NewPassword string `json:"newPassword"`
}

type ArticleCreateParam struct {
    Title string `json:"title"`
    Intro string `json:"intro"`
    Content string `json:"content"`
}

type ArticleRecord struct {
    Article
    FrontUser FrontUser `json:"frontUser"`
}

type FrontUserCreateParam struct {
    Username string `json:"username"`
    Email string `json:"email"`
}
