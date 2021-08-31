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
func Login(p *models.ParamLogin) (*models.User, error) {
	var user models.User
	user.UserName = p.UserName
	err := mysql.CheckUserPassword(p, &user)
	if err != nil {
		return nil, err
	}

	// 生成 JWT
	token, err := jwt.GenToken(user.UserId)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return &user, err
}

func GetAuthName(uid int64) (string, error) {
	return mysql.GetAuthName(uid)
}
