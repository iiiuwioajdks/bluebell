package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": nil,
		"msg":  code.Msg(),
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": nil,
		"msg":  msg,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"data": data,
		"msg":  CodeSuccess.Msg(),
	})
}
