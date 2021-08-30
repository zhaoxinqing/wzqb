package config

import (
	"Kilroy/app/model"
	"context"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ConfigPath = "config/conf.yaml"

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

//  获取配置信息
func GetConfig() (conf *Yaml) {
	yamlFile, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		panic("配置文件：读取失败")
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic("配置文件：解析失败")
	}
	return
}

// 初始化数据库
func InitDB() {
	conf := GetConfig()
	dsn := fmt.Sprintf("host= %v port= %v dbname= %v user=postgres password=gorm sslmode=disable TimeZone=Asia/Shanghai",
		conf.Database.Host, conf.Database.Port, conf.Database.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库：配置连接失败")
	}
	db.Debug()
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

var client *redis.Client
var ctx = context.Background()
var instance *Redis

func InitRedis() {
	conf := GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Database.Host,
		Password: conf.Database.Port,
		DB:       0,
		PoolSize: 132,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("Redis 连接失败 【错误】：%s", err.Error())
		return
	}
	fmt.Printf("Redis 连接成功： %s", pong)
}
