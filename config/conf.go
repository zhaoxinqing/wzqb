package config

import (
	"Kilroy/app/model"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	ConfigPath = "config/conf.yaml"
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
	conf, err := getConfig(ConfigPath)
	dsn := fmt.Sprintf("host= %v port= %v dbname= %v user=postgres password=gorm sslmode=disable TimeZone=Asia/Shanghai",
		conf.Database.Host, conf.Database.Port, conf.Database.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库：配置连接失败")
	}
	// err = autoMigrate(db)
	// if err != nil {
	// 	err = errors.New("数据库表：自动迁移失败")
	// 	fmt.Println(err)
	// }
	model.DB = db
}

//  获取配置信息
func getConfig(configPath string) (conf *Yaml, err error) {
	// Read configuration file
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("配置文件：读取失败")
	}
	// Parse the configuration file
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println("配置文件：解析失败")
	}
	return
}

// 数据库表自动迁移
func autoMigrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(new(model.User))
	err = db.AutoMigrate(new(model.UserInfo))
	return
}
