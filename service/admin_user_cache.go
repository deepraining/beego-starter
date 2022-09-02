package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/cache"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
    "strconv"
    "strings"
    "time"
)

const (
    redisDatabase = "sbsAdmin"
    redisExpire = 86400 * time.Second // 24 hours
    redisKeyAdmin = "admin"
    redisKeyResourceList = "resourceList"
)

// 删除用户缓存
func DelAdminUserCache(userId int64)  {
    err, adminUser := GetAdminUserItem(userId)
    if err != nil {
        logs.Error(err)
        return
    }
    key := redisDatabase + ":" + redisKeyResourceList + ":" + adminUser.Username
    err = cache.Delete(key)
    if err != nil {
        logs.Error(err)
    }
}

// 删除资源列表缓存
func DelAdminResourceListCache(userId int64)  {
    key := redisDatabase + ":" + redisKeyResourceList + ":" + strconv.FormatInt(userId, 10)
    err := cache.Delete(key)
    if err != nil {
        logs.Error(err)
    }
}

// 通过角色删除资源列表缓存
func DelAdminResourceListByRoleCache(roleId int64)  {
    var list []models.AdminUserRoleRelation
    utils.GetDB().Where("role_id = ?", roleId).Find(&list)

    keyPrefix := redisDatabase + ":" + redisKeyResourceList + ":"
    if len(list) > 0 {
        for _, item := range list {
            key := keyPrefix + strconv.FormatInt(item.UserId, 10)
            err := cache.Delete(key)
            if err != nil {
                logs.Error(err)
            }
        }
    }
}

// 通过角色删除资源列表缓存
func DelAdminResourceListByRoleIdsCache(roleIds *[]int64)  {
    var list []models.AdminUserRoleRelation
    utils.GetDB().Where("roleId in ?", roleIds).Find(&list)

    keyPrefix := redisDatabase + ":" + redisKeyResourceList + ":"
    if len(list) > 0 {
        for _, item := range list {
            key := keyPrefix + strconv.FormatInt(item.UserId, 10)
            err := cache.Delete(key)
            if err != nil {
                logs.Error(err)
            }
        }
    }
}

// 通过资源删除资源列表缓存
func DelAdminResourceListByResourceCache(resourceId int64)  {
    var list []models.AdminUserRoleRelation
    result := utils.GetDB().Raw(`
SELECT
  DISTINCT ar.user_id
FROM
  admin_role_resource_relation rr
  LEFT JOIN admin_user_role_relation ar ON rr.role_id = ar.role_id
WHERE rr.resource_id=?
`, resourceId).Scan(&list)

    if result.Error != nil {
        logs.Error(result.Error)
        return
    }

    keyPrefix := redisDatabase + ":" + redisKeyResourceList + ":"
    if len(list) > 0 {
        for _, item := range list{
            key := keyPrefix + strconv.FormatInt(item.UserId, 10)
            err := cache.Delete(key)
            if err != nil {
                logs.Error(err)
            }
        }
    }
}

// 获取用户缓存
func GetAdminUserCache(username string) *models.AdminUser {
    key := redisDatabase + ":" + redisKeyAdmin + ":" + username
    // 不能用空指针接收值
    var adminUser models.AdminUser
    err := cache.Get(key, &adminUser)
    if err != nil {
        // 缓存不存在的错误信息：cache does not exist
        if !strings.Contains(err.Error(), "not exist") {
            logs.Error(err)
        }
        return nil
    }
    if adminUser.Id > 0 {
        return &adminUser
    }
    return nil
}

// 设置用户缓存
func SetAdminUserCache(adminUser *models.AdminUser)  {
    key := redisDatabase + ":" + redisKeyAdmin + ":" + adminUser.Username
    err := cache.Set(key, adminUser, redisExpire)
    if err != nil {
        logs.Error(err)
    }
}

// 获取用户资源列表缓存
func GetAdminResourceListCache(userId int64) *[]models.AdminResource {
    key := redisDatabase + ":" + redisKeyResourceList + ":" + strconv.FormatInt(userId, 10)
    var result []models.AdminResource
    err := cache.Get(key, &result)
    if err != nil {
        // 缓存不存在的错误信息：cache does not exist
        if !strings.Contains(err.Error(), "not exist") {
            logs.Error(err)
        }
        return nil
    }
    if len(result) > 0 {
        return &result
    }
    return nil
}

// 设置用户资源列表缓存
func SetAdminResourceListCache(userId int64, resourceList *[]models.AdminResource) {
    key := redisDatabase + ":" + redisKeyResourceList + ":" + strconv.FormatInt(userId, 10)
    err := cache.Set(key, resourceList, redisExpire)
    if err != nil {
        logs.Error(err)
    }
}

