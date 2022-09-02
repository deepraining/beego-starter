package service

import (
    "errors"
    "github.com/beego/beego/v2/core/logs"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/utils"
    "github.com/jinzhu/copier"
    "time"
)

// 通过用户名获取用户信息
func GetAdminUserByUsername(username string) (error, *models.AdminUser) {
    adminUser := GetAdminUserCache(username)
    if adminUser != nil {
        return nil, adminUser
    }
    adminUser = &models.AdminUser{}
    result := utils.GetDB().Find(adminUser, "username = ?", username)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    if adminUser.Id > 0 && adminUser.Status == 1 {
        SetAdminUserCache(adminUser)
        return nil, adminUser
    }
    return nil, nil
}

// 注册用户
func AdminRegister(param *models.AdminUserParam) (error, *models.AdminUser) {
    if param.Username == "" {
        return errors.New(utils.CustomMsgPrefix+"用户名不能为空"), nil
    }
    if param.Password == "" {
        return errors.New(utils.CustomMsgPrefix+"密码不能为空"), nil
    }

    adminUser := &models.AdminUser{}
    copier.Copy(adminUser, param)
    adminUser.Status = 1

    // 查询是否有相同用户名的用户
    var count int64 = 0
    result := utils.GetDB().Model(adminUser).Where("username = ?", param.Username).Count(&count)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    if count > 0 {
        return errors.New(utils.CustomMsgPrefix+"用户名已存在"), nil
    }

    // 将密码进行加密操作
    encryptedPassword := utils.EncryptPassword(adminUser.Password)
    adminUser.Password = encryptedPassword
    result = utils.GetDB().Create(adminUser)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, adminUser
}

// 用户登录
func AdminLogin(username string, password string) (error, string) {
    err, adminUserDetails := LoadAdminUserByUsername(username)
    if err != nil {
        logs.Error(err)
        return err, ""
    }
    if !utils.ComparePassword(password, adminUserDetails.Password) {
        return errors.New(utils.CustomMsgPrefix+"密码不正确"), ""
    }
    err, token := utils.GenerateToken(username)
    if err != nil {
        logs.Error(err)
        return err, ""
    }
    insertAdminLoginLog(username)

    adminUser := &models.AdminUser{
        LastLoginTime: time.Now(),
    }
    result := utils.GetDB().Where("username = ?", username).Updates(adminUser)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, ""
    }

    return nil, token
}

// 添加登录记录
func insertAdminLoginLog(username string)  {
    _, user := GetAdminUserByUsername(username)
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
func AdminRoleListByUserId(userId int64) (error, *[]models.AdminRole) {
    list := &[]models.AdminRole{}
    result := utils.GetDB().Raw(`
select r.*
from admin_user_role_relation ar left join admin_role r on ar.role_id = r.id
where ar.user_id = ?
`, userId).Scan(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

// 通过用户ID获取资源列表
func AdminResourceListByUserId(userId int64) (error, *[]models.AdminResource) {
    list := GetAdminResourceListCache(userId)
    if list != nil {
        return nil, list
    }
    list = &[]models.AdminResource{}
    result := utils.GetDB().Raw(`
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

    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    if len(*list) != 0 {
        SetAdminResourceListCache(userId, list)
    }

    return nil, list
}

// 通过用户ID获取菜单列表
func AdminMenuListByUserId(userId int64) (error, *[]models.AdminMenu) {
    list := &[]models.AdminMenu{}
    result := utils.GetDB().Raw(`
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
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

// 获取用户信息
func LoadAdminUserByUsername(username string) (error, *models.AdminUserDetails) {
    // 获取用户信息
    err, adminUser := GetAdminUserByUsername(username)
    if err != nil {
        logs.Error(err)
        return err, nil
    }
    if adminUser != nil {
        adminUserDetails := &models.AdminUserDetails{}
        copier.Copy(adminUserDetails, adminUser)
        err, resourceList := AdminResourceListByUserId(adminUser.Id)
        if err != nil {
            logs.Error(err)
            return err, nil
        }
        adminUserDetails.ResourceList = resourceList
        return nil, adminUserDetails
    }
    return errors.New(utils.CustomMsgPrefix+"用户名或密码错误"), nil
}

// 刷新token
func AdminRefreshToken(oldToken string) (error, string) {
    return utils.RefreshHeadToken(oldToken)
}

// 获取用户项
func GetAdminUserItem(id int64) (error, *models.AdminUser) {
    adminUser := &models.AdminUser{}
    result := utils.GetDB().Find(adminUser, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    if adminUser.Id > 0 {
        return nil, adminUser
    }
    return nil, nil
}

// 用户列表
func AdminUserList(keyword string, pageSize int64, pageNum int64) (error, *[]models.AdminUser, int64) {
    limit := pageSize
    offset := pageSize * (pageNum - 1)

    query := utils.GetDB().Model(&models.AdminUser{})
    if keyword != "" {
        query.Where("username like ?", "%"+ keyword +"%").Or("nickname like ?", "%"+ keyword +"%")
    }
    list := &[]models.AdminUser{}
    var total int64 = 0
    query.Count(&total)
    result := query.Limit(int(limit)).Offset(int(offset)).Find(list)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil, 0
    }
    return nil, list, total
}

// 更新用户信息
func UpdateAdminUser(id int64, adminUser *models.AdminUser) (error, int64) {
    adminUser.Id = id
    rawUser := &models.AdminUser{}
    result := utils.GetDB().Find(rawUser, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    if rawUser.Password == adminUser.Password {
        // 与原加密密码相同的不需要修改
        adminUser.Password = ""
    } else if adminUser.Password != "" {
        adminUser.Password = utils.EncryptPassword(adminUser.Password)
    }

    // 账户名不能更改
    adminUser.Username = ""
    result = utils.GetDB().Updates(adminUser)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    DelAdminUserCache(id)
    return nil, result.RowsAffected
}

// 通过map更新(struct不更新0值)
func UpdateAdminUserByMap(id int64, adminUser *map[string]interface{}) (error, int64) {
    result := utils.GetDB().Model(&models.AdminUser{}).Where("id = ?", id).Updates(adminUser)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    DelAdminUserCache(id)
    return nil, result.RowsAffected
}

// 删除用户
func DeleteAdminUser(id int64) (error, int64) {
    DelAdminUserCache(id)
    result := utils.GetDB().Delete(&models.AdminUser{}, id)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    DeleteAdminResource(id)
    return nil, result.RowsAffected
}

// 更新用户的角色
func UpdateAdminUserRole(userId int64, roleIds *[]int64) (error, int64) {
    // 先删除原来的关系
    result := utils.GetDB().Where("user_id = ?", userId).Delete(&models.AdminUserRoleRelation{})
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    if roleIds != nil {
        // 建立新关系
        relationList := []models.AdminUserRoleRelation{}
        for _, item := range *roleIds{
            relationList = append(relationList, models.AdminUserRoleRelation{
                UserId: userId,
                RoleId: item,
            })
        }
        result = utils.GetDB().Create(relationList)
        if result.Error != nil {
            logs.Error(result.Error)
            return result.Error, 0
        }
        return nil, result.RowsAffected
    }
    return nil, 0
}

// 更新用户权限
func UpdateAdminUserPermission(userId int64, permissionIds *[]int64) (error, int64) {
    // 删除原所有权限关系
    result := utils.GetDB().Where("user_id = ?", userId).Delete(&models.AdminUserPermissionRelation{})
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }

    if permissionIds == nil || len(*permissionIds) == 0 {
        return nil, 0
    }

    // 获取用户所有角色权限
    rolePermissions := &[]models.AdminPermission{}
    result = utils.GetDB().Raw(`
select p.*
from admin_user_role_relation ar left join admin_role r on ar.role_id = r.id
  left join admin_role_permission_relation rp on r.id = rp.role_id
  left join admin_permission p on rp.permission_id=p.id
  where ar.user_id = ? and p.id is not null
`, userId).Scan(rolePermissions)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }

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
    result = utils.GetDB().Create(relationList)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    return nil, result.RowsAffected
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
func GetAdminPermissionList(userId int64) (error, *[]models.AdminPermission) {
    list := &[]models.AdminPermission{}
    result := utils.GetDB().Raw(`
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
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, nil
    }
    return nil, list
}

func UpdateAdminPassword(param *models.UpdateAdminUserPasswordParam) (error, int64) {
    if param.Username == "" || param.OldPassword == "" || param.NewPassword == "" {
        return nil, -1
    }

    adminUser := &models.AdminUser{}
    result := utils.GetDB().Where("username = ?", param.Username).Find(adminUser)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    if adminUser.Id == 0 {
        return nil, -2
    }

    if !utils.ComparePassword(param.OldPassword, adminUser.Password) {
        return nil, -3
    }

    updateUser := &models.AdminUser {
        Id: adminUser.Id,
        Password: utils.EncryptPassword(param.NewPassword),
    }

    result = utils.GetDB().Updates(updateUser)
    if result.Error != nil {
        logs.Error(result.Error)
        return result.Error, 0
    }
    DelAdminUserCache(adminUser.Id)
    return nil, 1
}
