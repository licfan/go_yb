package initialize

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"go_yb/global"
)

func Redis() *redis.Client {
	redisCfg := global.YB_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.YB_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.YB_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
	return nil
}
