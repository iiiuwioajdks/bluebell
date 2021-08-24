package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	// 注册 zap 日志路由
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(400, "pong")
	})

	// 用户模块
	userController := controller.NewUserController()
	{
		r.POST("/signup", userController.SignUpHandler)
		r.POST("/login", userController.LoginHandler)
	}
	return r
}
