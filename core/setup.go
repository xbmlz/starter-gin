package core

import (
	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/initialize"
)

// 项目配置
func Setup() {
	// init config
	initialize.Viper()

	// init logger
	initialize.Zap()

	if global.Config.Redis.Enable {
		// init redis
		initialize.Redis()
	} else {
		// init local cache
		initialize.BigCache()
	}

	// init datasource
	initialize.Gorm()
}
