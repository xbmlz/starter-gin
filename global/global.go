package global

import (
	"github.com/allegro/bigcache/v3"
	"github.com/go-redis/redis/v8"
	"github.com/xbmlz/starter-gin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config     config.Config      // 全局配置
	Log        *zap.Logger        // 全局日志
	DB         *gorm.DB           // 数据库实例
	Redis      *redis.Client      // redis实例
	LocalCache *bigcache.BigCache // 本地缓存
)
