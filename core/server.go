package core

import (
	"fmt"
	"go_yb/global"
	"go_yb/initialize"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"time"
	//"go.uber.org/zap"
	//"go_yb/global"
	//"time"
)

type server interface {
	//ListenAndServe() error
}

func RunWindowsServer() {
	//配置文件
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(yamlFile, &global.YB_CONFIG)
	if err != nil {
		fmt.Println(err)
	}

	//路由
	Router := initialize.Routers()

	//启动
	address := fmt.Sprintf(":%d", global.YB_CONFIG.System.Addr)
	s := http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	fmt.Println("启动成功YB", address)
	fmt.Println(s.ListenAndServe().Error())
}
