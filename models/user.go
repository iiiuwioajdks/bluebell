package models

type User struct {
	UserName string `db:"username"`
	Password string `db:"password"`
	UserId   int64  `db:"user_id"`
	Token    string
}
