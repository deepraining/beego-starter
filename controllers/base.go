package controllers

import (
    "encoding/json"
    "github.com/beego/beego/v2/core/logs"
    "github.com/beego/beego/v2/server/web"
    "github.com/deepraining/beego-starter/models"
    "github.com/deepraining/beego-starter/service"
    "github.com/deepraining/beego-starter/utils"
    "io"
    "strings"
)

var secureIgnoreUrls []string

func init()  {
    secureIgnoreUrlsStr, _ := web.AppConfig.String("secureIgnoreUrls")
    secureIgnoreUrls = strings.Split(secureIgnoreUrlsStr, ",")
}

type BaseController struct {
    web.Controller
    Username string
}

// 响应JSON数据
func (c *BaseController) JsonResult(data interface{}) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        logs.Error(err)
    }

    c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
    c.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
    _, err = io.WriteString(c.Ctx.ResponseWriter, string(jsonData))
    if err != nil {
        logs.Error(err)
    }

    c.StopRun()
}

// 响应失败，终止执行
func (c *BaseController) ApiFail(message string) {
    c.JsonResult(models.FailedResultWithMessage(message))
}

// 响应成功
func (c *BaseController) ApiSucceed(data interface{}) {
    c.JsonResult(models.SuccessResult(data))
}

func (c *BaseController) ApiSucceedWithMessage(data interface{}, message string) {
    c.JsonResult(models.SuccessResultWithMessage(data, message))
}

// 预校验
func (c *BaseController) Prepare() {
    requestPath := c.Ctx.Request.URL.Path
    // 如果是非 / 页，且末尾有/的，去掉末尾的/
    if len(requestPath) > 1 && strings.HasSuffix(requestPath, "/") {
        requestPath = requestPath[0:len(requestPath)-1]
    }
    // 校验登录状态
    jwtToken := c.Ctx.Request.Header.Get(utils.JwtTokenHeaderKey)
    err, username := utils.GetUserNameFromToken(jwtToken)
    if err != nil {
        c.ApiFail(utils.NormalizeErrorMessage(err))
    }
    c.Username = username
    // 是否是安全的URL链接
    isSecureUrl := utils.MatchUrl(secureIgnoreUrls, requestPath)

    if !isSecureUrl {
        // 未登录，校验安全链接
        if username == "" {
            // 响应未授权信息
            c.JsonResult(models.UnauthorizedResult())
        }
        // 校验权限
        err, adminUserDetails := service.LoadAdminUserByUsername(username)
        if err != nil {
            c.ApiFail(utils.NormalizeErrorMessage(err))
        }
        if adminUserDetails == nil || adminUserDetails.ResourceList == nil {
            // 响应未授权信息
            c.JsonResult(models.UnauthorizedResult())
        }

        resourceUrls := []string{}
        for _, item := range *adminUserDetails.ResourceList{
            resourceUrls = append(resourceUrls, item.Url)
        }

        // 已匹配上权限URL
        matched := utils.MatchUrl(resourceUrls, requestPath)
        if !matched {
            // 响应无权限信息
            c.JsonResult(models.ForbiddenResult())
        }
    }
}
