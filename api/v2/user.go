package v2

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_yb/global"
	"go_yb/model/request"
	"go_yb/utils"
)

func Register(c *gin.Context) {

}

func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBind(&L)
	err := utils.Verify(&L)
	global.YB_LOG.Info("s", zap.Any("aa", err))
}
