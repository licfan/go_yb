package main

import (
	"go_yb/core"
	"go_yb/global"
	"go_yb/initialize"
)

func main() {
	core.InitConfig()
	global.YB_LOG = core.Zap()
	global.YB_REDIS = initialize.Redis()
	global.YB_DB = initialize.Gorm()
	db,_ := global.YB_DB.DB()
	defer db.Close()

	core.RunWindowsServer()
}
