package controller

import (
	"bluebell/middleware"
	"errors"
	"github.com/gin-gonic/gin"
)

/**
用来获取 *gin.Context 里面存的东西
*/

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 用来拿中间件 auth 中存进去的 userId
func GetCurrentUser(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(middleware.UserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
