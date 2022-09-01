package utils

import (
    "errors"
    "fmt"
    "github.com/beego/beego/v2/core/logs"
    "github.com/beego/beego/v2/server/web"
    "github.com/golang-jwt/jwt/v4"
    "strings"
    "time"
)

var TokenHeaderKey string
var secret string
var expiration int64
var TokenHead string

func init()  {
    TokenHeaderKey, _ = web.AppConfig.String("jwtTokenHeader")
    secret, _ = web.AppConfig.String("jwtSecret")
    expiration, _ = web.AppConfig.Int64("jwtExpiration")
    TokenHead, _ = web.AppConfig.String("jwtTokenHead")
}

// 根据用户名生成JWT的token
func GenerateToken(username string) (error, string) {
    claims :=  &jwt.RegisteredClaims{
        Subject: username,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiration) * time.Second)),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
    tokenStr, err := token.SignedString([]byte(secret))
    if err != nil {
        logs.Error(err)
        return err, ""
    }
    return nil, tokenStr
}

// 从token中获取登录用户名
func GetUserNameFromToken(tokenStr string) (error, string)  {
    tokenStr = strings.Trim(tokenStr, " ")
    if tokenStr == "" {
        return nil, ""
    }
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        logs.Error(err)
        return err, ""
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return nil, fmt.Sprintf("%v", claims["sub"])
    } else {
        logs.Error("JWT格式验证失败:", tokenStr)
        return errors.New("Bad jwt token format"), ""
    }
}

// 当原来的token没过期时是可以刷新的
// @param oldToken 带tokenHead的token
func RefreshHeadToken(oldToken string) (error, string) {
    if oldToken == "" {
        return nil, ""
    }
    tokenStr := oldToken[len(TokenHead):]
    if tokenStr == "" {
        return nil, ""
    }
    err, username := GetUserNameFromToken(tokenStr)
    if err != nil {
        return err, ""
    }
    // token校验不通过，或者已失效
    if username == "" {
        return nil, ""
    }
    // 生成新的token
    return GenerateToken(username)
}
