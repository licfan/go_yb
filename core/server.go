package core

import (
	"fmt"
	"go_yb/global"
	"go_yb/initialize"
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
