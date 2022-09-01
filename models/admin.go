package models

import "time"

type AdminLoginLog struct {
    Id  int64 `gorm:"default:(-)"`
    UserId  int64 `gorm:"default:(-)"`
    Ip  string `gorm:"default:(-)"` // ip地址
    Address  string `gorm:"default:(-)"` // 地址
    UserAgent  string `gorm:"default:(-)"` // 浏览器登录类型
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminMenu struct {
    Id  int64 `gorm:"default:(-)"`
    ParentId  int64 `gorm:"default:(-)"` // 父级ID
    Title  string `gorm:"default:(-)"` // 菜单名称
    Level int32 `gorm:"default:(-)"` // 菜单级数
    Sort int32 `gorm:"default:(-)"` // 菜单排序
    Name  string `gorm:"default:(-)"` // 前端名称
    Icon  string `gorm:"default:(-)"` // 前端图标
    Hidden int32 `gorm:"default:(-)"` // 前端隐藏
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminPermission struct {
    Id  int64 `gorm:"default:(-)"`
    ParentId  int64 `gorm:"default:(-)"` // 父级权限id
    Name  string `gorm:"default:(-)"` // 名称
    Value  string `gorm:"default:(-)"` // 权限值
    Icon  string `gorm:"default:(-)"` // 图标
    Type int32 `gorm:"default:(-)"` // 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）
    Uri  string `gorm:"default:(-)"` // 前端资源路径
    Status int32 `gorm:"default:(-)"` // 启用状态；0->禁用；1->启用
    Sort int32 `gorm:"default:(-)"` // 排序
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminResource struct {
    Id  int64 `gorm:"default:(-)"`
    Name  string `gorm:"default:(-)"` // 资源名称
    Url  string `gorm:"default:(-)"` // 资源URL
    Description  string `gorm:"default:(-)"` // 描述
    CategoryId  int64 `gorm:"default:(-)"` // 资源分类ID
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminResourceCategory struct {
    Id  int64 `gorm:"default:(-)"`
    Name  string `gorm:"default:(-)"` // 分类名称
    Sort int32 `gorm:"default:(-)"` // 排序
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminRole struct {
    Id  int64 `gorm:"default:(-)"`
    Name  string `gorm:"default:(-)"` // 名称
    Description  string `gorm:"default:(-)"` // 描述
    UserCount int32 `gorm:"default:(-)"` // 后台用户数量
    Status int32 `gorm:"default:(-)"` // 启用状态：0->禁用；1->启用
    Sort int32 `gorm:"default:(-)"` // 排序
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminRoleMenuRelation struct {
    Id  int64 `gorm:"default:(-)"`
    RoleId  int64 `gorm:"default:(-)"` // 角色ID
    MenuId  int64 `gorm:"default:(-)"` // 菜单ID
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminRolePermissionRelation struct {
    Id  int64 `gorm:"default:(-)"`
    RoleId  int64 `gorm:"default:(-)"` // 角色ID
    PermissionId  int64 `gorm:"default:(-)"` // 权限ID
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminRoleResourceRelation struct {
    Id  int64 `gorm:"default:(-)"`
    RoleId  int64 `gorm:"default:(-)"` // 角色ID
    ResourceId  int64 `gorm:"default:(-)"` // 资源ID
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminUser struct {
    Id  int64 `gorm:"default:(-)"`
    Username  string `gorm:"default:(-)"` // 用户名
    Password  string `gorm:"default:(-)"` // 加密密码
    Avatar  string `gorm:"default:(-)"` // 头像
    Email  string `gorm:"default:(-)"` // 邮箱
    Nickname  string `gorm:"default:(-)"` // 昵称
    Note  string `gorm:"default:(-)"` // 备注信息
    LastLoginTime time.Time `gorm:"default:(-)"` // 最后登录时间
    Status int32 `gorm:"default:(-)"` // 启用状态：0->禁用；1->启用
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminUserPermissionRelation struct {
    Id  int64 `gorm:"default:(-)"`
    UserId  int64 `gorm:"default:(-)"` // 用户ID
    PermissionId  int64 `gorm:"default:(-)"` // 权限ID
    Type int32 `gorm:"default:(-)"` // 类型：1->增加权限，-1->减少权限
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}

type AdminUserRoleRelation struct {
    Id  int64 `gorm:"default:(-)"`
    UserId  int64 `gorm:"default:(-)"` // 用户ID
    RoleId  int64 `gorm:"default:(-)"` // 角色ID
    CreateTime time.Time `gorm:"default:(-)"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)"` // 更新时间
}
