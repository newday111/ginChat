package redisdb

import (
	"context"
	"go_jichu/conf"
	"go_jichu/internal/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

var GlobalRdb redis.UniversalClient

func InitRedisConn(cfg *conf.ConfigStruct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	isCluster := cfg.REDIS.IsCluster // 用于切换集群或者单机使用
	if isCluster {
		GlobalRdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    cfg.REDIS.RedisAddr,
			Password: cfg.REDIS.Password, // 如果有密码
		})
		utils.AccessLog.Info("正在以【集群模式】链接Redis....")
	} else {
		GlobalRdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    cfg.REDIS.RedisAddr,
			Password: cfg.REDIS.Password,
			DB:       cfg.REDIS.RedisDb,
		})
		utils.AccessLog.Info("正在以【单机模式】链接Redis....")
	}

	if err := GlobalRdb.Ping(ctx).Err(); err != nil {
		return err
	}
	utils.AccessLog.Info("redis connection start success!!!")
	return nil
}
