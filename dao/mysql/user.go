package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
)

const secret = "dweqfojjidvfushorgijfcnvierrwog"

var (
	ErrorNP        = errors.New("用户名或密码错误")
	ErrorUserExist = errors.New("用户已存在")
)

func CheckUserPassword(p *models.ParamLogin, u *models.User) error {
	sqlStr := "select user_id,username,password from user where username=?"
	err := db.Get(u, sqlStr, p.UserName)
	// 用户不存在判断
	if err == sql.ErrNoRows {
		// 一般不直接告诉用户名不存在，防止攻击
		return ErrorNP
	}
	if err != nil {
		return err
	}
	p.Password = encryptPassword(p.Password)
	if u.Password == p.Password {
		return nil
	}
	return ErrorNP
}

func CheckUserExist(username string) error {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	err := db.Get(&count, sqlStr, username)
	if count > 0 {
		return ErrorUserExist
	}
	return err
}

func InsertUser(u *models.User) error {
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	u.Password = encryptPassword(u.Password)
	// 执行 sql 语句
	_, err := db.Exec(sqlStr, u.UserId, u.UserName, u.Password)
	fmt.Println(err)
	return err
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
