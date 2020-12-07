package main

import (
	//"go_yb/global"
	"go_yb/core"
	//"go_yb/global"
)

func main() {
	//global.YB_REDIS = ""
	//global.YB_LOG = ""
	//global.YB_gorm = ""
	core.RunWindowsServer()
}
