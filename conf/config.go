package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"wzqb/models"

	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

// GetConfigInformation 获取配置信息
func Run() (err error) {
	// read config
	conf, err := confRead("./conf/config.yaml")
	if err != nil {
		return
	}
	// db
	err = initDB(conf.DSN)
	if err != nil {
		return errors.New("数据库连接失败")
	}
	// redis
	err = initRedis(conf.Redis)
	if err != nil {
		return errors.New("Redis连接失败")
	}

	return
}

// ConfRead ...
func confRead(path string) (conf ConfigInfos, err error) {
	// read
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// unmarshal
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

// DB ...

// initDB ...
func initDB(dsn string) (err error) {
	//
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	//
	DB.AutoMigrate(&models.AccountTable{})
	//
	return
}

func initRedis(info RedisConfig) (err error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     info.Host, //
		Password: info.Pwd,  //
		DB:       info.DB,   //
	})
	_, err = Redis.Ping().Result()
	if err != nil {
		return
	}
	return
}
