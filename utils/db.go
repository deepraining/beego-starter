package utils

import "gorm.io/gorm"

var db *gorm.DB

func InitDB(dbInstance *gorm.DB)  {
    db = dbInstance
}

func GetDB() *gorm.DB  {
    return db
}
