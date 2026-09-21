package redisdb

import (
	"context"
	"go_jichu/internal/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

var GlobalRdb redis.UniversalClient

func InitRedisConn() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	isCluster := false // 用于切换集群或者单机使用
	if isCluster {
		GlobalRdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: []string{
				"127.0.0.1:7000",
				"127.0.0.1:7001",
				"127.0.0.1:7002",
			},
			Password: "your_password", // 如果有密码
		})
		utils.AccessLog.Info("正在以【集群模式】链接Redis....")
	} else {
		GlobalRdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "123456",
			DB:       0,
		})
		utils.AccessLog.Info("正在以【单机模式】链接Redis....")
	}

	if err := GlobalRdb.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}
