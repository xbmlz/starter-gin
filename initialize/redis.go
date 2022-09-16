package initialize

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/xbmlz/starter-gin/global"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.Config.Redis
	redisAddr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisConfig.Password,
		DB:       redisConfig.Database,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Error("Redis connect failed, err:", zap.Error(err))
	} else {
		global.Log.Sugar().Infof("Redis is connected. ")
		global.Redis = client
	}
}
