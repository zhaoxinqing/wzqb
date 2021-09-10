package config

import (
	"Kilroy/app/model"
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB ...
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

// InitRedis ...
func InitRedis() {
	conf := GetConfig()
	client := redis.NewClient(&redis.Options{
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

// autoMigrate ...
func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(new(model.User))
	db.AutoMigrate(new(model.UserInfo))
}
