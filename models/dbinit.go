package models

import (
	"Kilroy/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB 初始化数据库配置
func InitDB(config *config.Config) {
	var gdb *gorm.DB
	// gdb, err = gorm.Open(config.Gorm.DBType, config.Gorm.DSN)
	gdb, err := gorm.Open(postgres.Open(config.Gorm.DSN))
	if err != nil {
		panic(err)
	}
	//  默认使用单数表命名
	// gdb.SingularTable(true)
	// if config.Gorm.Debug {
	// 	gdb.LogMode(true)
	// 	// gdb.SetLogger(log.New(os.Stdout, "\r\n", 0))
	// }
	// gdb.DB().SetMaxIdleConns(config.Gorm.MaxIdleConns)
	// gdb.DB().SetMaxOpenConns(config.Gorm.MaxOpenConns)
	// gdb.DB().SetConnMaxLifetime(time.Duration(config.Gorm.MaxLifetime) * time.Second)
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
