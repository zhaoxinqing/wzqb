package ctrl

import (
	"Kilroy/app/common"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Test struct{}

type Yaml struct {
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
}

// 获取
func (i Test) Yaml(c *gin.Context) {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
	common.ResSuccess(c, nil)
}
