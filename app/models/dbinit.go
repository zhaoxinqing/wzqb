package models

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

//profile variables
type Conf struct {
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

func (c *Conf) GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("config/conf.yaml")
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
func InitDB(config *Conf) {
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
