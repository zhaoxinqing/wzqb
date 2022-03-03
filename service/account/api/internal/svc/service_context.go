package svc

import (
	"go-template/service/account/api/internal/config"
	"go-template/service/account/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//
	dns := c.DataSourceName
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Println(err)
	}
	DBAutoMigrate(db)
	//
	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}

//  autoMigrate ...
func DBAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.AccountTable{}) // 账户
	db.AutoMigrate(&model.RoleTable{})    // 角色
}
