package initialize

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	var Router = gin.Default()
	return Router
}
