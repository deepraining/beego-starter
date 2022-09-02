package models

import "time"

type AdminLoginLog struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    UserId  int64 `gorm:"default:(-)" json:"userId"`
    Ip  string `gorm:"default:(-)" json:"ip"` // ip地址
    Address  string `gorm:"default:(-)" json:"address"` // 地址
    UserAgent  string `gorm:"default:(-)" json:"userAgent"` // 浏览器登录类型
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminMenu struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    ParentId  int64 `gorm:"default:(-)" json:"parentId"` // 父级ID
    Title  string `gorm:"default:(-)" json:"title"` // 菜单名称
    Level int32 `gorm:"default:(-)" json:"level"` // 菜单级数
    Sort int32 `gorm:"default:(-)" json:"sort"` // 菜单排序
    Name  string `gorm:"default:(-)" json:"name"` // 前端名称
    Icon  string `gorm:"default:(-)" json:"icon"` // 前端图标
    Hidden int32 `gorm:"default:(-)" json:"hidden"` // 前端隐藏
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminPermission struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    ParentId  int64 `gorm:"default:(-)" json:"parentId"` // 父级权限id
    Name  string `gorm:"default:(-)" json:"name"` // 名称
    Value  string `gorm:"default:(-)" json:"value"` // 权限值
    Icon  string `gorm:"default:(-)" json:"icon"` // 图标
    Type int32 `gorm:"default:(-)" json:"type"` // 权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）
    Uri  string `gorm:"default:(-)" json:"uri"` // 前端资源路径
    Status int32 `gorm:"default:(-)" json:"status"` // 启用状态；0->禁用；1->启用
    Sort int32 `gorm:"default:(-)" json:"sort"` // 排序
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminResource struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    Name  string `gorm:"default:(-)" json:"name"` // 资源名称
    Url  string `gorm:"default:(-)" json:"url"` // 资源URL
    Description  string `gorm:"default:(-)" json:"description"` // 描述
    CategoryId  int64 `gorm:"default:(-)" json:"categoryId"` // 资源分类ID
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminResourceCategory struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    Name  string `gorm:"default:(-)" json:"name"` // 分类名称
    Sort int32 `gorm:"default:(-)" json:"sort"` // 排序
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminRole struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    Name  string `gorm:"default:(-)" json:"name"` // 名称
    Description  string `gorm:"default:(-)" json:"description"` // 描述
    UserCount int32 `gorm:"default:(-)" json:"userCount"` // 后台用户数量
    Status int32 `gorm:"default:(-)" json:"status"` // 启用状态：0->禁用；1->启用
    Sort int32 `gorm:"default:(-)" json:"sort"` // 排序
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminRoleMenuRelation struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    RoleId  int64 `gorm:"default:(-)" json:"roleId"` // 角色ID
    MenuId  int64 `gorm:"default:(-)" json:"menuId"` // 菜单ID
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminRolePermissionRelation struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    RoleId  int64 `gorm:"default:(-)" json:"roleId"` // 角色ID
    PermissionId  int64 `gorm:"default:(-)" json:"permissionId"` // 权限ID
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminRoleResourceRelation struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    RoleId  int64 `gorm:"default:(-)" json:"roleId"` // 角色ID
    ResourceId  int64 `gorm:"default:(-)" json:"resourceId"` // 资源ID
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminUser struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    Username  string `gorm:"default:(-)" json:"username"` // 用户名
    Password  string `gorm:"default:(-)" json:"password"` // 加密密码
    Avatar  string `gorm:"default:(-)" json:"avatar"` // 头像
    Email  string `gorm:"default:(-)" json:"email"` // 邮箱
    Nickname  string `gorm:"default:(-)" json:"nickname"` // 昵称
    Note  string `gorm:"default:(-)" json:"note"` // 备注信息
    LastLoginTime time.Time `gorm:"default:(-)" json:"lastLoginTime"` // 最后登录时间
    Status int32 `gorm:"default:(-)" json:"status"` // 启用状态：0->禁用；1->启用
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminUserPermissionRelation struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    UserId  int64 `gorm:"default:(-)" json:"userId"` // 用户ID
    PermissionId  int64 `gorm:"default:(-)" json:"permissionId"` // 权限ID
    Type int32 `gorm:"default:(-)" json:"type"` // 类型：1->增加权限，-1->减少权限
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}

type AdminUserRoleRelation struct {
    Id  int64 `gorm:"default:(-)" json:"id"`
    UserId  int64 `gorm:"default:(-)" json:"userId"` // 用户ID
    RoleId  int64 `gorm:"default:(-)" json:"roleId"` // 角色ID
    CreateTime time.Time `gorm:"default:(-)" json:"createTime"` // 创建时间
    UpdateTime time.Time `gorm:"default:(-)" json:"updateTime"` // 更新时间
}
