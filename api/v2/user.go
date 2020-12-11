package v2

import (
	"github.com/gin-gonic/gin"
	"go_yb/model/request"
	"go_yb/model/response"
	"go_yb/utils"
)

func Register(c *gin.Context) {

}

func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBind(&L)
	validateErr := utils.Verify(&L)
	if validateErr != "" {
		response.ErrorMsgResponse(validateErr,c)
		return
	}

}
