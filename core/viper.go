package core

import (
	"fmt"
	"go_yb/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/**
*	初始化配置文件
 */
func InitConfig() {
	//配置文件
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &global.YB_CONFIG)
	if err != nil {
		fmt.Println(err)
	}
}
