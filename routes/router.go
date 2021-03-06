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
	v1 := r.Group("")
	userController := controller.NewUserController()
	{
		// 注册
		v1.POST("/signup", userController.SignUpHandler)
		// 登录
		v1.POST("/login", userController.LoginHandler)
	}

	v1.Use(middleware.JWTAuthMiddleware())
	communityController := controller.NewCommunityController()
	postController := controller.NewPostController()
	{
		v1.GET("/community", communityController.CommunityHandler)
		v1.GET("/community/:id", communityController.CommunityDetailHandler)

		v1.POST("/post", postController.CreatePost)
		v1.GET("/post/:id", postController.ShowPost)
		v1.GET("/posts", postController.GetPostList)
		v1.GET("/posts2", postController.GetPostList2)
		v1.POST("/vote", postController.Vote)
	}

	return r
}
