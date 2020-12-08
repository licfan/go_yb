package main

import (
	_ "fmt"
	"go_yb/core"
	"go_yb/global"
	//"go_yb/global"
)

func main() {

	//global.YB_REDIS = ""
	global.YB_LOG = core.Zap()
	//global.YB_gorm = ""
	core.RunWindowsServer()
}
