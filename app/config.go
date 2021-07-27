package app

import (
	"Kilroy/app/model"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// yaml配置文件
type Yaml struct {
	Database Postgres `yaml:"postgres"`
	Redis    Redis    `yaml:"redis"`
}

// 数据库
type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

// 缓存库
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// 初始化数据库
func InitDB() {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=postgres password=gorm sslmode=disable TimeZone=Asia/Shanghai",
		conf.Database.Host, conf.Database.Port, conf.Database.DBName)
	// log.Printf("dsn: %v ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(new(model.Users))
	err = db.AutoMigrate(new(model.UserInfo))
	if err != nil {
		fmt.Println(err)
	}
	model.DB = db
}
