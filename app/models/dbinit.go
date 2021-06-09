package models

import (
	"Kilroy/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

// InitDB 初始化数据库配置
func InitDB(config *config.Conf) {
	dsn := "host=localhost user=postgres password=gorm dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

// Migration 初始化数据库表
func Migration() {
	var err error
	err = DB.AutoMigrate(new(Users))
	if err != nil {
		fmt.Println(err)
	}

}
