package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/response"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IUser interface {
	SignUpHandler(c *gin.Context)
}

type User struct {
	DB *sqlx.DB
}

func NewUserController() IUser {
	return User{DB: mysql.GetDB()}
}

func (u User) SignUpHandler(c *gin.Context) {
	// 1.获取参数和校验
	var p models.ParamSignUp
	// 通过 binding 自动校验，参数不合法就返回错误
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("Sign Up with invalid param", zap.Error(err))
		response.Fail(c, nil, "请求参数有误--err:"+err.Error())
		return
	}

	// 2.业务处理
	err = logic.SignUp(&p)
	if err != nil {
		zap.L().Error("logic sign up error", zap.Error(err))
		response.Fail(c, nil, err.Error())
		return
	}
	// 3.响应
	response.Success(c, nil, "sign up success")
}
