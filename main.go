package main

import (
	"fmt"
	"go-template/wzqb"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func main() {
	// config
	conf := GetConfigInformation("config.yanml")
	fmt.Println(conf)
	// Creates a router without any middleware by default
	r := gin.Default()

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", wzqb.GetAllMenu)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/api")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use()
	{
		authorized.POST("/login", wzqb.Login)

		// // nested group
		// testing := authorized.Group("testing")
		// visit 0.0.0.0:8080/testing/analytics
		// testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// ConfigInformation config info
type ConfigInformation struct {
	DB struct {
		Host   string `yaml:"host"`
		User   string `yaml:"user"`
		Pwd    string `yaml:"pwd"`
		DBname string `yaml:"dbname"`
	} `yaml:"db"`
}

// GetConfigInformation 获取配置信息
func GetConfigInformation(configPath string) (conf ConfigInformation) {
	// read
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	// unmarshal
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
