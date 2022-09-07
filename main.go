package main

import (
    "encoding/json"
    "fmt"
    beegoCache "github.com/beego/beego/v2/client/cache"
    _ "github.com/beego/beego/v2/client/cache/redis"
    "github.com/beego/beego/v2/core/logs"
    "github.com/beego/beego/v2/server/web"
    "github.com/bwmarrin/snowflake"
    "github.com/deepraining/beego-starter/cache"
    _ "github.com/deepraining/beego-starter/routers"
    "github.com/deepraining/beego-starter/utils"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/schema"
    "os"
    "path/filepath"
)

func main() {
    initDB()
    initCache()
    initLogger()
    initUuid()

    web.Run()
}

func initDB() {
    host, _ := web.AppConfig.String("mysqlhost")
    database, _ := web.AppConfig.String("mysqldb")
    username, _ := web.AppConfig.String("mysqluser")
    password, _ := web.AppConfig.String("mysqlpass")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true,
        },
    })

    if err != nil {
        logs.Error("初始化数据库失败:", err)
        os.Exit(1)
    }
    utils.InitDB(db)
}

func initCache() {
    var redisConfig struct {
        Conn     string `json:"conn"`
        Password string `json:"password"`
        DbNum    string `json:"dbNum"`
    }

    redisConfig.DbNum = web.AppConfig.DefaultString("redisdb", "0")
    redisConfig.Conn, _ = web.AppConfig.String("redishost")
    redisConfig.Password, _ = web.AppConfig.String("redispass")

    bc, err := json.Marshal(&redisConfig)
    if err != nil {
        logs.Error("初始化Redis缓存失败:", err)
        os.Exit(1)
    }
    redisCache, err := beegoCache.NewCache("redis", string(bc))

    if err != nil {
        logs.Error("初始化Redis缓存失败:", err)
        os.Exit(1)
    }

    cache.Init(redisCache)
}

func initLogger() {
    logs.SetLogFuncCall(true)
    _ = logs.SetLogger("console")

    //logs.Async(1e3)

    logPath := filepath.Join("logs", "log.log")

    if _, err := os.Stat("logs"); os.IsNotExist(err) {
        _ = os.MkdirAll("logs", 0755)
    }

    config := make(map[string]interface{}, 1)

    config["filename"] = logPath
    config["perm"] = "0755"
    config["rotate"] = true
    config["maxLines"] = 1000000
    config["maxsize"] = 1<<28
    config["daily"] = true
    config["maxdays"] = 30

    if level := web.AppConfig.DefaultString("logLevel", "Trace"); level != "" {
        switch level {
        case "Emergency":
            config["level"] = logs.LevelEmergency
        case "Alert":
            config["level"] = logs.LevelAlert
        case "Critical":
            config["level"] = logs.LevelCritical
        case "Error":
            config["level"] = logs.LevelError
        case "Warning":
            config["level"] = logs.LevelWarning
        case "Notice":
            config["level"] = logs.LevelNotice
        case "Informational":
            config["level"] = logs.LevelInformational
        case "Debug":
            config["level"] = logs.LevelDebug
        }
    }
    b, err := json.Marshal(config)
    if err != nil {
        logs.Error("初始化文件日志时出错:", err)
        _ = logs.SetLogger("file", `{"filename":"`+logPath+`"}`)
    } else {
        _ = logs.SetLogger(logs.AdapterFile, string(b))
    }
}

func initUuid()  {
    uuidNode, _ := web.AppConfig.Int64("uuidNode")
    node, err := snowflake.NewNode(uuidNode)
    if err != nil {
        logs.Error("初始化uuid生成器:", err)
        os.Exit(1)
    }
    utils.InitUuid(node)
}
