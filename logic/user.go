package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	// 判断用户是否存在
	err := mysql.CheckUserExist(p.UserName)
	if err != nil {
		// 数据库查询出错
		return err
	}
	// 生成 UID
	userId := snowflake.GetId()

	// 构造 user 实例
	u := models.User{
		UserName: p.UserName,
		UserId:   userId,
		Password: p.Password,
	}
	// 存入数据库
	err = mysql.InsertUser(&u)
	return err
}

// Login 返回token和error
func Login(p *models.ParamLogin) (string, error) {
	var user models.User
	err := mysql.CheckUserPassword(p, &user)
	if err != nil {
		return "", err
	}

	// 生成 JWT
	return jwt.GenToken(user.UserId)

}
