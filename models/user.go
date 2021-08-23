package models

type User struct {
	UserName string `db:"username"`
	UserId   int64  `db:"user_id"`
	Password string `db:"password"`
}
