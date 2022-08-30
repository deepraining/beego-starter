package models

import "time"

type AdminLoginLog struct {
    Id int64
    UserId int64
    Ip string // ip地址
    Address string // 地址
    UserAgent string // 浏览器登录类型
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminMenu struct {
    Id int64
    ParentId int64 // 父级ID
    Title string // 菜单名称
    Level int32 // 菜单级数
    Sort int32 // 菜单排序
    Name string // 前端名称
    Icon string // 前端图标
    Hidden int32 // 前端隐藏
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminPermission struct {
    Id int64
    ParentId int64 // 父级权限id
    Name string // 名称
    Value string // 权限值
    Icon string // 图标
    Type int32 // 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）
    Uri string // 前端资源路径
    Status int32 // 启用状态；0->禁用；1->启用
    Sort int32 // 排序
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminResource struct {
    Id int64
    Name string // 资源名称
    Url string // 资源URL
    Description string // 描述
    CategoryId int64 // 资源分类ID
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminResourceCategory struct {
    Id int64
    Name string // 分类名称
    Sort int32 // 排序
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminRole struct {
    Id int64
    Name string // 名称
    Description string // 描述
    UserCount int32 // 后台用户数量
    Status int32 // 启用状态：0->禁用；1->启用
    Sort int32 // 排序
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminRoleMenuRelation struct {
    Id int64
    RoleId int64 // 角色ID
    MenuId int64 // 菜单ID
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminRolePermissionRelation struct {
    Id int64
    RoleId int64 // 角色ID
    PermissionId int64 // 权限ID
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminRoleResourceRelation struct {
    Id int64
    RoleId int64 // 角色ID
    ResourceId int64 // 资源ID
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminUser struct {
    Id int64
    Username string // 用户名
    Password string // 加密密码
    Avatar string // 头像
    Email string // 邮箱
    Nickname string // 昵称
    Note string // 备注信息
    LastLoginTime time.Time // 最后登录时间
    Status int32 // 启用状态：0->禁用；1->启用
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminUserPermissionRelation struct {
    Id int64
    UserId int64 // 用户ID
    PermissionId int64 // 权限ID
    Type int32 // 类型：1->增加权限，-1->减少权限
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}

type AdminUserRoleRelation struct {
    Id int64
    UserId int64 // 用户ID
    RoleId int64 // 角色ID
    CreateTime time.Time // 创建时间
    UpdateTime time.Time // 更新时间
}
