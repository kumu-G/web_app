package redis

import (
	"context"
	"fmt"
	"go.uber.org/zap"

	"web_app/setting"

	"github.com/go-redis/redis/v8"
)

// 实际生产环境下 context.Background() 按需替换

var (
	client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	zap.L().Info("init redis success")
	return nil
}

func Close() {
	_ = client.Close()
}
