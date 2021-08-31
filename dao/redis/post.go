package redis

import (
	"bluebell/models"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func CreatePost(p *models.Post) error {
	ctx := context.Background()
	// 创建这两个要么同时成功，要么同时失败，所以用事务
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(ctx, KeyPostTimeZSet, &redis.Z{Score: float64(time.Now().Unix()), Member: p.ID}).Result()

	pipeline.ZAdd(ctx, KeyPostScoreZSet, &redis.Z{Score: float64(0), Member: p.ID}).Result()

	_, err := pipeline.Exec(ctx)
	return err
}
