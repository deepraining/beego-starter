package utils

import (
    "github.com/beego/beego/v2/core/logs"
    "golang.org/x/crypto/bcrypt"
    "reflect"
    "strconv"
    "strings"
)

func NormalizeUrl(url string) string {
    // 如果是非 / 页，且末尾有/的，去掉末尾的/
    if len(url) > 1 && strings.HasSuffix(url, "/") {
        return url[0:len(url)-1]
    }
    return url
}

// 匹配URL
func MatchUrl(urls []string, path string) bool {
    path = NormalizeUrl(path)
    for _, item := range urls{
        url := NormalizeUrl(item)
        if !strings.HasSuffix(url, "*") {
            // 不以*结尾，当做普通的对待
            if url == path {
                return true
            }
        } else if strings.HasSuffix(url, "/**/*") {
            // /**/* 匹配子孙目录
            if strings.HasPrefix(path, url[0:len(url)-5]) {
                return true
            }

        } else if strings.HasSuffix(url, "/*") {
            // /* 匹配子目录
            if len(strings.Split(url, "/")) == len(strings.Split(path, "/")) && strings.HasPrefix(path, url[0:len(url)-2]) {
                return true
            }
        }
    }

    return false
}

// 复制结构体属性(父->子)
func CopyStructFields(source interface{}, target interface{})  {
    sourceVal := reflect.ValueOf(source)
    targetVal := reflect.ValueOf(target)
    fieldsNum := sourceVal.NumField()
    for i := 0; i < fieldsNum; i++ {
        sourceField := sourceVal.Field(i)
        targetField := targetVal.Field(i)
        targetField.Set(sourceField.Elem())
    }
}

// 加密密码
func EncryptPassword(password string) string {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
    if err != nil {
        logs.Error(err)
    }
    return string(hash)
}

// 验证密码
func ComparePassword(password string, encryptedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(password), []byte(encryptedPassword))
    if err != nil {
        logs.Error(err)
        return false
    }
    return true
}

// 字符转数字
func StringToInt64(val string, defaults int64) int64 {
    if val == "" {
        return defaults
    }
    intVal, err := strconv.Atoi(val)
    if err != nil {
        return defaults
    }
    return int64(intVal)
}
