package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	// 注册 zap 日志路由
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	// 用户模块
	userController := controller.NewUserController()
	{
		// 注册
		r.POST("/signup", userController.SignUpHandler)
		// 登录
		r.POST("/login", userController.LoginHandler)
	}

	return r
}
