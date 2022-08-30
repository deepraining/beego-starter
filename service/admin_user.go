package service

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/beego/beego/v2/core/validation"
    "github.com/senntyou/beego-starter/models"
    "github.com/senntyou/beego-starter/utils"
    "time"
)

// 通过用户名获取用户信息
func GetAdminUserByUsername(username string) *models.AdminUser {
    adminUser := GetAdminUserCache(username)
    if adminUser != nil {
        return adminUser
    }
    adminUser = &models.AdminUser{}
    utils.GetDB().First(adminUser, "username = ?", username)
    if adminUser.Id > 0 && adminUser.Status == 1 {
        SetAdminUserCache(adminUser)
        return adminUser
    }
    return nil
}

// 注册用户
func AdminRegister(param *models.AdminUserParam) *models.AdminUser {
    valid := validation.Validation{}
    b, err := valid.Valid(param)
    if err != nil {
        logs.Error(err)
    }
    if !b {
        // validation does not pass
        for _, err := range valid.Errors {
            panic(err)
        }
    }

    adminUser := &models.AdminUser{}
    utils.CopyStructFields(param, adminUser)
    adminUser.Status = 1

    // 查询是否有相同用户名的用户
    var count int64 = 0
    utils.GetDB().Where("username = ?", param.Username).Count(&count)
    if count > 0 {
        return nil
    }

    // 将密码进行加密操作
    encryptedPassword := utils.EncryptPassword(adminUser.Password)
    adminUser.Password = encryptedPassword
    utils.GetDB().Create(adminUser)
    return adminUser
}

// 用户登录
func AdminLogin(username string, password string) string {
    adminUserDetails := LoadAdminUserByUsername(username)
    if !utils.ComparePassword(password, adminUserDetails.Password) {
        panic("密码不正确")
    }
    token := utils.GenerateToken(username)
    insertAdminLoginLog(username)
    updateAdminLoginTimeByUsername(username)
    return token
}

// 根据用户名修改登录时间
func updateAdminLoginTimeByUsername(username string)  {
    adminUser := &models.AdminUser{
        LastLoginTime: time.Now(),
    }
    utils.GetDB().Where("username = ?", username).Updates(adminUser)
}

// 添加登录记录
func insertAdminLoginLog(username string)  {
    user := GetAdminUserByUsername(username)
    if user == nil {
        return
    }

    loginLog := &models.AdminLoginLog{
        UserId: user.Id,
        Ip: "",
    }
    utils.GetDB().Create(loginLog)
}

// 通过用户ID获取角色列表
func AdminRoleListByUserId(userId int64) *[]models.AdminRole {
    list := &[]models.AdminRole{}
    utils.GetDB().Raw(`
select r.*
from admin_user_role_relation ar left join admin_role r on ar.role_id = r.id
where ar.user_id = ?
`, userId).Scan(list)
    return list
}

// 通过用户ID获取资源列表
func AdminResourceListByUserId(userId int64) *[]models.AdminResource {
    list := GetAdminResourceListCache(userId)
    if list != nil {
        return list
    }

    list = &[]models.AdminResource{}
    utils.GetDB().Raw(`
SELECT
  ur.*
FROM
  admin_user_role_relation ar
LEFT JOIN admin_role r ON ar.role_id = r.id
LEFT JOIN admin_role_resource_relation rrr ON r.id = rrr.role_id
LEFT JOIN admin_resource ur ON ur.id = rrr.resource_id
WHERE
  ar.user_id = ?
AND ur.id IS NOT NULL
GROUP BY
  ur.id
`, userId).Scan(list)

    if len(*list) != 0 {
        SetAdminResourceListCache(userId, list)
    }

    return list
}

// 通过用户ID获取菜单列表
func AdminMenuListByUserId(userId int64) *[]models.AdminMenu {
    list := &[]models.AdminMenu{}
    utils.GetDB().Raw(`
SELECT
  m.*
FROM
  admin_user_role_relation arr
    LEFT JOIN admin_role r ON arr.role_id = r.id
    LEFT JOIN admin_role_menu_relation rmr ON r.id = rmr.role_id
    LEFT JOIN admin_menu m ON rmr.menu_id = m.id
WHERE
  arr.user_id = ?
    AND m.id IS NOT NULL
GROUP BY
  m.id
`, userId).Scan(list)
    return list
}

// 获取用户信息
func LoadAdminUserByUsername(username string) *models.AdminUserDetails {
    // 获取用户信息
    adminUser := GetAdminUserByUsername(username)
    if adminUser != nil {
        adminUserDetails := &models.AdminUserDetails{}
        utils.CopyStructFields(adminUser, adminUserDetails)
        resourceList := AdminResourceListByUserId(adminUser.Id)
        adminUserDetails.ResourceList = resourceList
        return adminUserDetails
    }
    panic("用户名或密码错误")
    return nil
}

// 刷新token
func AdminRefreshToken(oldToken string) string {
    return utils.RefreshHeadToken(oldToken)
}

// 获取用户项
func GetAdminUserItem(id int64) *models.AdminUser {
    adminUser := &models.AdminUser{}
    utils.GetDB().First(adminUser, id)
    if adminUser.Id > 0 {
        return adminUser
    }
    return nil
}

// 用户列表
func AdminUserList(keyword string, pageSize int64, pageNum int64) (*[]models.AdminUser, int64) {
    limit := pageSize
    offset := pageSize * (pageNum - 1)

    query := utils.GetDB()
    if keyword != "" {
        query.Where("username like ?", "%"+ keyword +"%").Or("nickname like ?", "%"+ keyword +"%")
    }
    list := &[]models.AdminUser{}
    var total int64 = 0
    query.Count(&total)
    query.Limit(int(limit)).Offset(int(offset)).Find(list)
    return list, total
}

// 更新用户信息
func UpdateAdminUser(id int64, adminUser *models.AdminUser) int64 {
    adminUser.Id = id
    rawUser := &models.AdminUser{}
    utils.GetDB().First(rawUser, id)
    if rawUser.Password == adminUser.Password {
        // 与原加密密码相同的不需要修改
        adminUser.Password = ""
    } else if adminUser.Password != "" {
        adminUser.Password = utils.EncryptPassword(adminUser.Password)
    }

    // 账户名不能更改
    adminUser.Username = ""
    result := utils.GetDB().Updates(adminUser)
    DelAdminUserCache(id)
    return result.RowsAffected
}

// 删除用户
func DeleteAdminUser(id int64) int64 {
    DelAdminUserCache(id)
    result := utils.GetDB().Delete(&models.AdminUser{Id: id})
    DeleteAdminResource(id)
    return result.RowsAffected
}

// 更新用户的角色
func UpdateAdminUserRole(userId int64, roleIds *[]int64) int64 {
    // 先删除原来的关系
    utils.GetDB().Delete(&models.AdminUserRoleRelation{UserId: userId})
    if roleIds != nil {
        // 建立新关系
        relationList := []models.AdminUserRoleRelation{}
        for _, item := range *roleIds{
            relationList = append(relationList, models.AdminUserRoleRelation{
                UserId: userId,
                RoleId: item,
            })
        }
        result := utils.GetDB().Create(relationList)
        return result.RowsAffected
    }
    return 0
}

// 更新用户权限
func UpdateAdminUserPermission(userId int64, permissionIds *[]int64) int64 {
    // 删除原所有权限关系
    utils.GetDB().Delete(&models.AdminUserPermissionRelation{UserId: userId})

    if permissionIds == nil || len(*permissionIds) == 0 {
        return 0
    }

    // 获取用户所有角色权限
    rolePermissions := &[]models.AdminPermission{}
    utils.GetDB().Raw(`
select p.*
from admin_user_role_relation ar left join admin_role r on ar.role_id = r.id
  left join admin_role_permission_relation rp on r.id = rp.role_id
  left join admin_permission p on rp.permission_id=p.id
  where ar.user_id = ? and p.id is not null
`, userId).Scan(rolePermissions)

    rolePermissionIds := []int64{}
    for _, item := range *rolePermissions {
        rolePermissionIds = append(rolePermissionIds, item.Id)
    }

    // 筛选出+权限，角色权限里没有，就加上
    addPermissionIdList := []int64{}
    for _, item:=range *permissionIds{
        found := false
        for _, item2 := range rolePermissionIds{
            if item == item2 {
                found = true
                break
            }
        }
        if !found {
            addPermissionIdList = append(addPermissionIdList, item)
        }
    }

    // 筛选出-权限，角色权限里有，但permissionIds没有，则去掉
    subPermissionIdList := []int64{}
    for _, item:=range rolePermissionIds{
        found := false
        for _, item2 := range *permissionIds{
            if item == item2 {
                found = true
                break
            }
        }
        if !found {
            subPermissionIdList = append(addPermissionIdList, item)
        }
    }

    relationList := []models.AdminUserPermissionRelation{}
    relationList = append(relationList, *convertAdminPermissionRelation(userId, 1, &addPermissionIdList)...)
    relationList = append(relationList, *convertAdminPermissionRelation(userId, 1, &subPermissionIdList)...)
    result := utils.GetDB().Create(relationList)
    return result.RowsAffected
}

func convertAdminPermissionRelation(userId int64, type_ int32, permissionIdList *[]int64) *[]models.AdminUserPermissionRelation {
    relationList := []models.AdminUserPermissionRelation{}
    for _, item := range *permissionIdList {
        relation := models.AdminUserPermissionRelation{
            UserId: userId,
            Type: type_,
            PermissionId: item,
        }
        relationList = append(relationList, relation)
    }
    return &relationList
}

// 获取权限列表
func GetAdminPermissionList(userId int64) *[]models.AdminPermission {
    list := &[]models.AdminPermission{}
    utils.GetDB().Raw(`
SELECT
      p.*
    FROM
      admin_user_role_relation ar
      LEFT JOIN admin_role r ON ar.role_id = r.id
      LEFT JOIN admin_role_permission_relation rp ON r.id = rp.role_id
      LEFT JOIN admin_permission p ON rp.permission_id = p.id
    WHERE
      ar.user_id = ?
      AND p.id IS NOT NULL
      AND p.id NOT IN (
        SELECT
          p.id
        FROM
          admin_user_permission_relation pr
          LEFT JOIN admin_permission p ON pr.permission_id = p.id
        WHERE
          pr.type = - 1
          AND pr.user_id = ?
      )
    UNION
    SELECT
      p.*
    FROM
      admin_user_permission_relation pr
      LEFT JOIN admin_permission p ON pr.permission_id = p.id
    WHERE
      pr.type = 1
      AND pr.user_id = ?
`, userId).Scan(list)
    return list
}

func UpdateAdminPassword(param *models.UpdateAdminUserPasswordParam) int64 {
    valid := validation.Validation{}
    b, err := valid.Valid(param)
    if err != nil {
        logs.Error(err)
    }
    if !b {
        // validation does not pass
        for _, err := range valid.Errors {
            panic(err)
        }
    }

    if param.Username == "" || param.OldPassword == "" || param.NewPassword == "" {
        return -1
    }

    adminUser := &models.AdminUser{}
    utils.GetDB().Where("username = ?", param.Username).First(adminUser)
    if adminUser.Id == 0 {
        return -2
    }

    if !utils.ComparePassword(param.OldPassword, adminUser.Password) {
        return -3
    }

    updateUser := &models.AdminUser {
        Id: adminUser.Id,
        Password: utils.EncryptPassword(param.NewPassword),
    }

    utils.GetDB().Updates(updateUser)
    DelAdminUserCache(adminUser.Id)
    return 1
}
