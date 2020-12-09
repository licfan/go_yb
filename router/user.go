package router

import (
	"github.com/gin-gonic/gin"
	v2 "go_yb/api/v2"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", v2.Login)
	}
}
