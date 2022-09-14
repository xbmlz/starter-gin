package global

import (
	"github.com/xbmlz/starter-gin/config"
	"go.uber.org/zap"
)

var (
	Config config.Config // 全局配置
	Logger *zap.Logger   // 全局日志
)
