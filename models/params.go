package models

// 定义请求参数的结构体

type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}
