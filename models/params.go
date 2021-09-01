package models

// 定义请求参数的结构体

type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamPostList struct {
	CommunityId int64  `json:"community_id" form:"community_id"` // 可以为空，空的时候查全部，有的时候就只查社区
	Page        int64  `form:"page"`
	Size        int64  `form:"size"`
	Order       string `form:"order"`
}
