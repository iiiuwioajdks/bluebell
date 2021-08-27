package controller

import (
	"bluebell/logic"
	"bluebell/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// 跟社区相关

type ICommunity interface {
	CommunityHandler(c *gin.Context)
	CommunityDetailHandler(c *gin.Context)
}

type Community struct {
}

func (com Community) CommunityDetailHandler(c *gin.Context) {
	idstr := c.Param("id")
	communityId, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	data, err := logic.GetCommunityDetail(communityId)
	if err != nil {
		zap.L().Error("GetCommunityDetail failed", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}
	// 返回数据
	response.Success(c, data)
}

func NewCommunityController() ICommunity {
	return Community{}
}

func (com Community) CommunityHandler(c *gin.Context) {

	// 查询数据(id,name)
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("GetCommunityList failed", zap.Error(err))
		response.ResponseError(c, response.CodeServerBusy)
		return
	}
	// 返回数据
	response.Success(c, data)
}
