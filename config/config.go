package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 配置参数
type Config struct {
	Gorm     Gorm
	Postgres Postgres
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
	DSN          string
}

// Postgres 配置参数
type Postgres struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

var cc *Config

func GetConfig() *Config {
	return cc
}

// LoadConfig 加载配置
func LoadConfig(path string) (c *Config, err error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	mode := v.GetString("runMode")
	newpath := fmt.Sprintf("./config/%s.yaml", mode)
	v.SetConfigFile(newpath)
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	c = &Config{}
	cc = c

	c.Postgres.Host = v.GetString("postgres.host")
	c.Postgres.Port = v.GetInt("postgres.port")
	c.Postgres.User = v.GetString("postgres.user")
	c.Postgres.Password = v.GetString("postgres.password")
	c.Postgres.DBName = v.GetString("postgres.db_name")
	c.Postgres.Parameters = v.GetString("postgres.parameters")
	c.Gorm.Debug = v.GetBool("gorm.debug")
	c.Gorm.DBType = v.GetString("gorm.db_type")
	c.Gorm.MaxLifetime = v.GetInt("gorm.max_lifetime")
	c.Gorm.MaxOpenConns = v.GetInt("gorm.max_open_conns")
	c.Gorm.MaxIdleConns = v.GetInt("gorm.max_idle_conns")
	c.Gorm.TablePrefix = v.GetString("gorm.table_prefix")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		c.Postgres.Host, c.Postgres.User, c.Postgres.DBName, c.Postgres.Password, c.Postgres.Port)
	c.Gorm.DSN = dsn

	return
}
