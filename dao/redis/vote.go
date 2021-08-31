package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"math"
	"time"
)

/*
direction=1:
	之前没投过票0，现在赞成1			差值绝对值：1 +432
	之前投过反对票-1，现在改投赞成票	1   差值绝对值：2 +432*2
direction=0
	之前投过反对票-1，现在取消投票0  	差值绝对值：1 +432
	之前投过赞成票1，现在取消投票0	    差值绝对值：1 —432
direction=-1
	之前没投过票0，现在反对-1			差值绝对值：1 -432
	之前投过赞成票1，现在改投反对票-1	差值绝对值：2	-432*2
*/
const (
	ScorePerVote = 432
	OneWeekInS   = 7 * 24 * 3600
)

var ErrVoteTimeExpire = errors.New("投票时间已过")

func Vote(userId, postId string, direction float64) error {
	ctx := context.Background()
	// 1.判断投票限制
	// 取发布时间
	postTime := rdb.ZScore(ctx, KeyPostTimeZSet, postId).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInS {
		return ErrVoteTimeExpire
	}
	// 2和3里面的 分数增减以及记录用户数据 步骤需要用事务处理
	pipeline := rdb.TxPipeline()
	// 2.更新分数
	// 先查之前的投票记录
	ov := rdb.ZScore(ctx, KeyPostVotedZSet+postId, userId).Val()
	var dir float64
	// 现在的值大于之前的值，就表明要加分，看上面规律
	if direction > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - direction) // 计算两次投票的差值
	pipeline.ZIncrBy(ctx, KeyPostScoreZSet, dir*diff*ScorePerVote, postId)

	// 3.记录用户为该帖子投票的数据
	if direction == 0 { // 如果取消投票，就要删掉原来的值
		pipeline.ZRem(ctx, KeyPostVotedZSet+postId, userId)
	} else {
		pipeline.ZAdd(ctx, KeyPostVotedZSet+postId, &redis.Z{Score: direction, Member: userId})

	}
	_, err := pipeline.Exec(ctx)
	return err
}
