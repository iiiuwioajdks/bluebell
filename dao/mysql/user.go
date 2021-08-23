package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
)

const secret = "dweqfojjidvfushorgijfcnvierrwog"

func CheckUserExist(username string) error {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	err := db.Get(&count, sqlStr, username)
	if count > 0 {
		return errors.New("用户已存在")
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
