package config

import (
	"Kilroy/app/model"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 初始化数据库
func InitDB() {
	conf := GetConfig()
	dsn := fmt.Sprintf("host= %v port= %v dbname= %v user=postgres password=gorm sslmode=disable TimeZone=Asia/Shanghai",
		conf.Database.Host, conf.Database.Port, conf.Database.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库：配置连接失败")
	}
	autoMigrate(db)
	model.DB = db
}

// 数据库表自动迁移
func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(new(model.User))
	err = db.AutoMigrate(new(model.UserInfo))
	if err != nil {
		err = errors.New("数据库表：自动迁移失败")
		fmt.Println(err)
	}
	return
}
