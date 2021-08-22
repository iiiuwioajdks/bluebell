package routes

import (
	"github.com/gin-gonic/gin"
	"web_app/logger"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	// 注册 zap 日志路由
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(400, "pong")
	})

	return r
}
