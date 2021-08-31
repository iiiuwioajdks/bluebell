package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 投票功能
/*
1.用户投票的数据
2.用户投票算法：阮一峰(本项目使用简单的投票分数)
*/
// Vote 帖子投票

//投一票加432分,一天 时间戳 86400/200 -> 两百个赞成票可以加一天
/*
direction=1:
	之前没投过票，现在赞成			差值绝对值：1 +432
	之前投过反对票，现在改投赞成票	差值绝对值：2 +432*2
direction=0
	之前投过赞成票，现在取消投票	差值绝对值：1 —432
	之前投过反对票，现在取消投票	差值绝对值：1 +432
direction=-1
	之前没投过票，现在反对			差值绝对值：1 -432
	之前投过赞成票，现在改投反对票	差值绝对值：2	-432*2

投票限制：
	每个帖子自发表起一个星期内允许投票，超过一个星期就不允许了
	到期之后将redis中保存的赞成和反对票数存入MySQL中
	到期后删除 KeyPostVotedZSet
*/

func Vote(userId int64, v *models.VoteData) error {
	zap.L().Debug("VoteDorPost",
		zap.Int64("UserId", userId),
		zap.Int64("PostId", v.PostID),
		zap.Int8("direction", v.Direction))
	return redis.Vote(strconv.Itoa(int(userId)), strconv.Itoa(int(v.PostID)), float64(v.Direction))
}
