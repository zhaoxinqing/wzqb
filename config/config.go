package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

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
