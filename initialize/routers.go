package initialize

import (
	"github.com/gin-gonic/gin"
	"go_yb/router"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)//注册用户路由
	return Router
}
