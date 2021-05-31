package models

import (
	"Kilroy/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB 初始化数据库配置
func InitDB(config *config.Conf) {
	var gdb *gorm.DB
	gdb, err := gorm.Open(postgres.Open(config.DB.Host))
	if err != nil {
		panic(err)
	}
	DB = gdb
}

// DB ...
var DB *gorm.DB

// Migration 初始化数据库表
func Migration() {
	var err error
	// err = sys.DB.AutoMigrate(new(sys.Menu))
	// err = sys.DB.AutoMigrate(new(sys.Admins))
	// err = sys.DB.AutoMigrate(new(sys.RoleMenu))
	// err = sys.DB.AutoMigrate(new(sys.Role))
	// err = sys.DB.AutoMigrate(new(sys.AdminsRole))
	// err = sys.DB.AutoMigrate(new(sys.OperationLog))
	// err = sys.DB.AutoMigrate(new(sys.LoginLog))

	// err = sys.DB.AutoMigrate(new(model.PortraitGrid))
	if err != nil {
		fmt.Println(err)
	}

}
