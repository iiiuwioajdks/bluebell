package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IUser interface {
	SignUpHandler(c *gin.Context)
	LoginHandler(c *gin.Context)
}

type User struct {
	DB *sqlx.DB
}

func NewUserController() IUser {
	return User{DB: mysql.GetDB()}
}

func (u User) LoginHandler(c *gin.Context) {
	// 1.获取参数和校验
	var p models.ParamLogin
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	// 2.业务处理
	user, err := logic.Login(&p)
	if err != nil {
		zap.L().Error("logic login error", zap.Error(err))
		if errors.Is(err, mysql.ErrorNP) {
			response.ResponseError(c, response.CodeInvalidPassword)
			return
		}
		response.ResponseError(c, response.CodeServerBusy)
		return
	}
	response.Success(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserId),
		"user_name": user.UserName,
		"token":     user.Token,
	})
}

func (u User) SignUpHandler(c *gin.Context) {
	// 1.获取参数和校验
	var p models.ParamSignUp
	// 通过 binding 自动校验，参数不合法就返回错误
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("Sign Up with invalid param", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	// 2.业务处理
	err = logic.SignUp(&p)
	if err != nil {
		zap.L().Error("logic sign up error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			response.ResponseError(c, response.CodeUserExist)
			return
		}
		response.ResponseError(c, response.CodeServerBusy)
		return
	}
	// 3.响应
	response.Success(c, nil)
}
