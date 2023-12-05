package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go_web/settings"
)

var ctx = context.Background()
var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host,
			cfg.Port), //"localhost:6379"
		Password: cfg.Password, // 没有密码，默认值
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		//zap.L().Error("redis数据库连接失败", zap.Error(err))
		fmt.Println("rdb.Ping(ctx).Result():", err)
		return err
	} else {
		fmt.Println("redis数据库连接成功...")
	}
	return
}

func Close() {
	_ = rdb.Close()
}
