package redis

import (
	"bluebell/models"
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var (
	CkeyScore string
	CkeyTime  string
)

func CreatePost(p *models.Post) error {
	ctx := context.Background()
	// 创建这两个要么同时成功，要么同时失败，所以用事务
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(ctx, KeyPostTimeZSet, &redis.Z{Score: float64(time.Now().Unix()), Member: p.ID})

	pipeline.ZAdd(ctx, KeyPostScoreZSet, &redis.Z{Score: float64(0), Member: p.ID})

	// 把帖子 id 加入 社区的set
	CkeyScore = KeyCommunitySet + "score:" + strconv.Itoa(int(p.CommunityID))
	CkeyTime = KeyCommunitySet + "time:" + strconv.Itoa(int(p.CommunityID))
	pipeline.ZAdd(ctx, CkeyScore, &redis.Z{Score: float64(0), Member: p.ID})
	pipeline.ZAdd(ctx, CkeyTime, &redis.Z{Score: float64(time.Now().Unix()), Member: p.ID})

	_, err := pipeline.Exec(ctx)
	return err
}

func GetPostIDInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取ID
	ctx := context.Background()
	key := KeyPostScoreZSet
	if p.Order == "time" {
		key = KeyPostTimeZSet
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// zrevrange 从高到低
	result, err := rdb.ZRevRange(ctx, key, start, end).Result()
	return result, err
}

// GetPostVoteData 根据 ids 查询投票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64,0,len(ids))
	//ctx := context.Background()
	//for _,id := range ids {
	//	key := KeyPostVotedZSet+id
	//	// 查找 key 中 分数是 1 的元素的数量，-》统计每篇帖子赞成数的数量
	//	val := rdb.ZCount(ctx, key, "1", "1").Val()
	//	data = append(data,val)
	//}

	// 完全没必要一个发一次请求，可以使用 pipeline 一次性发送多条数据
	ctx := context.Background()
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := KeyPostVotedZSet + id
		pipeline.ZCount(ctx, key, "1", "1")
	}
	exec, err := pipeline.Exec(ctx)
	if err != nil {
		return
	}
	data = make([]int64, 0, len(exec))
	for _, cmder := range exec {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := p.Order
	if orderKey == "score" {
		orderKey = KeyCommunitySet + "score:" + strconv.Itoa(int(p.CommunityId))
	} else {
		orderKey = KeyCommunitySet + "time:" + strconv.Itoa(int(p.CommunityId))
	}

	ctx := context.Background()
	// 使用 ZInterstroe（取交集）  把分区的帖子set与帖子分数 zset 生成一个新的zset

	//cKey := KeyCommunitySet + strconv.Itoa(int(p.CommunityId))
	//key := orderKey + strconv.Itoa(int(p.CommunityId))
	//if rdb.Exists(ctx,key).Val() < 1 {
	//	// 不存在则需要计算
	//	// 获取
	//	resultCP, err := rdb.SMembers(ctx, cKey).Result()
	//	fmt.Println(resultCP)
	//	for i := range resultCP {
	//		score := rdb.ZScore(ctx,orderKey,resultCP[i]).Val()
	//		rdb.ZAdd(ctx,key,&redis.Z{Score: score,Member: resultCP[i]})
	//	}
	//	// 设置一个超时时间，以便在投票后重新统计
	//	rdb.Expire(ctx,key,6*time.Second)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	// 存在则直接查询

	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// zrevrange 从高到低
	result, err := rdb.ZRevRange(ctx, orderKey, start, end).Result()
	return result, err
}
