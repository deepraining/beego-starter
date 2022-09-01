package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/cache"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
    "strconv"
)

const (
    redisDatabase = "sbsAdmin"
    redisExpire = 86400 // 24 hours
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
    _ = cache.Delete(key)
}

// 删除资源列表缓存
func DelAdminResourceListCache(userId int64)  {
    key := redisDatabase + ":" + redisKeyResourceList + ":" + strconv.FormatInt(userId, 10)
    _ = cache.Delete(key)
}

// 通过角色删除资源列表缓存
func DelAdminResourceListByRoleCache(roleId int64)  {
    list := &[]models.AdminUserRoleRelation{}
    utils.GetDB().Where("roleId = ?", roleId).Find(list)

    keyPrefix := redisDatabase + ":" + redisKeyResourceList + ":"
    listLen := len(*list)
    if listLen > 0 {
        for _, item := range *list {
            key := keyPrefix + strconv.FormatInt(item.UserId, 10)
            _ = cache.Delete(key)
        }
    }
}

// 通过角色删除资源列表缓存
func DelAdminResourceListByRoleIdsCache(roleIds *[]int64)  {
    list := &[]models.AdminUserRoleRelation{}
    utils.GetDB().Where("roleId in ?", roleIds).Find(list)

    keyPrefix := redisDatabase + ":" + redisKeyResourceList + ":"
    listLen := len(*list)
    if listLen > 0 {
        for _, item := range *list {
            key := keyPrefix + strconv.FormatInt(item.UserId, 10)
            _ = cache.Delete(key)
        }
    }
}

// 通过资源删除资源列表缓存
func DelAdminResourceListByResourceCache(resourceId int64)  {
    list := &[]models.AdminUserRoleRelation{}
    utils.GetDB().Raw(`
SELECT
  DISTINCT ar.user_id
FROM
  admin_role_resource_relation rr
  LEFT JOIN admin_user_role_relation ar ON rr.role_id = ar.role_id
WHERE rr.resource_id=?
`, resourceId).Scan(list)

    keyPrefix := redisDatabase + ":" + redisKeyResourceList + ":"
    listLen := len(*list)
    if listLen > 0 {
        for _, item := range *list{
            key := keyPrefix + strconv.FormatInt(item.UserId, 10)
            _ = cache.Delete(key)
        }
    }
}

// 获取用户缓存
func GetAdminUserCache(username string) *models.AdminUser {
    key := redisDatabase + ":" + redisKeyAdmin + ":" + username
    adminUser := &models.AdminUser{}
    _ = cache.Get(key, adminUser)
    if adminUser.Id > 0 {
        return adminUser
    }
    return nil
}

// 设置用户缓存
func SetAdminUserCache(adminUser *models.AdminUser)  {
    key := redisDatabase + ":" + redisKeyAdmin + ":" + adminUser.Username
    _ = cache.Set(key, adminUser, redisExpire)
}

// 获取用户资源列表缓存
func GetAdminResourceListCache(userId int64) *[]models.AdminResource {
    key := redisDatabase + ":" + redisKeyResourceList + ":" + strconv.FormatInt(userId, 10)
    result := &[]models.AdminResource{}
    _ = cache.Get(key, result)
    return result
}

// 设置用户资源列表缓存
func SetAdminResourceListCache(userId int64, resourceList *[]models.AdminResource) {
    key := redisDatabase + ":" + redisKeyResourceList + ":" + strconv.FormatInt(userId, 10)
    _ = cache.Set(key, resourceList, redisExpire)
}

