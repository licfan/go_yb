package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkResponse(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func OkMsgResponse(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func ErrorResponse(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func ErrorMsgResponse(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
