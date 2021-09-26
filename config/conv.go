package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const ConfigPath = "./conf.yaml"

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

//  GetConfig ...
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
