package initialize

import (
	"admin/common"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitCache() {
	switch common.CONFIG.System.CacheType {
	case "redis":
		InitRedis()
	default:
		InitRedis()
	}

}

func InitRedis() {
	r := common.CONFIG.Redis
	common.CACHE = redis.NewClient(&redis.Options{
		Addr:     r.Addr,     // 使用默认地址
		Password: r.Password, // 未设置密码
		DB:       r.DB,       // use default DB
	})

	pong, err := common.CACHE.Ping().Result()
	if err != nil {
		common.LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		//[admin] 2022/05/28 - 17:03:33   info    redis connect ping response:    {"pong": "PONG"}
		common.LOG.Info("redis connect ping response:", zap.String("pong", pong))
	}

}
