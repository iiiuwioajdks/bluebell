package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

type IPost interface {
	CreatePost(c *gin.Context)
	ShowPost(c *gin.Context)
	GetPostList(c *gin.Context)
	Vote(c *gin.Context)
	GetPostList2(c *gin.Context)
}

type Post struct {
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// Vote 帖子投票功能
func (p Post) Vote(c *gin.Context) {
	// 参数校验
	v := new(models.VoteData)
	if err := c.ShouldBindJSON(v); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			response.ResponseErrorWithMsg(c, response.CodeInvalidParam, errs.Error())
			return
		}
		zap.L().Error("get vote data error", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	userId, err := GetCurrentUser(c)
	if err != nil {
		response.ResponseError(c, response.CodeNeedLogin)
		return
	}
	// 业务逻辑
	err = logic.Vote(userId, v)
	if err != nil {
		zap.L().Error("logic vote error", zap.Error(err))
		response.ResponseError(c, response.CodeServerBusy)
		return
	}
	response.Success(c, nil)
}

// GetPostList 获取帖子列表，并且分页
func (p Post) GetPostList(c *gin.Context) {
	// 获取分页参数
	// 第几页
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	// 一页多少数据
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	data, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("get post list error", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}
	response.Success(c, data)
}

// GetPostList2 按分数或时间排序帖子
func (p Post) GetPostList2(c *gin.Context) {
	// 获取分页参数
	// p.order 是 time 或 score ，代表按照时间查或者按照分数查
	paramPost := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: OrderTime,
	}

	err := c.ShouldBindQuery(paramPost)
	if err != nil {
		zap.L().Error("get post list error", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	data, err := logic.GetPostListNew(paramPost)

	if err != nil {
		zap.L().Error("get post list 2.0 error", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}
	response.Success(c, data)
}

// ShowPost 获得某一贴子的具体信息
func (post Post) ShowPost(c *gin.Context) {
	// 获取 id 参数
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get detail id error", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	// 传数据库
	var pd *models.PostDetail
	pd, err = logic.ShowPost(id)
	if err != nil {
		zap.L().Error("show post error ", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}
	response.Success(c, pd)
}

func NewPostController() IPost {
	return Post{}
}

// CreatePost 创建帖子
func (post Post) CreatePost(c *gin.Context) {
	// 获取参数
	var err error
	var p models.Post
	if err = c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("shouldBind error", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	// 从中间件 auth 中拿用户 ID,request 里面已经封装了 GetCurrentUser 方法了
	p.AuthorID, err = GetCurrentUser(c)
	if err != nil {
		zap.L().Error("get userid error", zap.Error(err))
		response.ResponseError(c, response.CodeNeedLogin)
		return
	}

	// 创建
	if err := logic.CreatePost(&p); err != nil {
		zap.L().Error("logic create post error", zap.Error(err))
		response.ResponseError(c, response.CodeServerBusy)
		return
	}

	response.Success(c, "帖子创建成功")

}
