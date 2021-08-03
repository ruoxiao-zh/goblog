package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"goblog/pkg/config"
	"goblog/pkg/logger"
)

var Rdb *redis.Client

// InitClient 初始化连接
func InitClient() *redis.Client {

	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetString("cache.redis.host") + ":" + config.GetString("cache.redis.port"),
		Password: config.GetString("cache.redis.password"), // no password set
		DB:       config.GetInt("cache.redis.database"),    // use default DB
		PoolSize: 100,                                            // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Rdb.Ping(ctx).Result()
	logger.LogError(err)

	return Rdb
}
