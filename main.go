package main

import (
	"Kilroy/app"
	"Kilroy/app/model"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var c ConfigFile
	conf := c.GetConf()
	InitDB(conf)
	router := gin.New()
	app.RegisterRouting(router.Group("/v1"))
	router.Run(":8990")
}

func InitDB(config *ConfigFile) {
	InitDBa(config) // 初始化数据库
	Migration()     // 数据库表迁移（自创建数据库）
}

// DB ...
var DB *gorm.DB

//profile variables
type ConfigFile struct {
	Database Postgres `yaml:"postgres"`
	Redis    Redis    `yaml:"redis"`
}
type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Driver   string `yaml:"driver"`
}

func (c *ConfigFile) GetConf() *ConfigFile {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

// InitDB 初始化数据库配置
func InitDBa(config *ConfigFile) {
	dsn := "host=localhost user=postgres password=gorm dbname=kilroy port=9900 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

// Migration 初始化数据库表
func Migration() {
	var err error
	err = DB.AutoMigrate(new(model.Users))
	err = DB.AutoMigrate(new(model.UserInfo))
	if err != nil {
		fmt.Println(err)
	}

}
