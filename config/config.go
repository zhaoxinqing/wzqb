package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//profile variables
type Conf struct {
	DB    DB    `yaml:"db"`
	Redis Redis `yaml:"redis"`
}
type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Driver   string `yaml:"driver"`
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
