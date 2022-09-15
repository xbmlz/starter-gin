package core

import "github.com/xbmlz/starter-gin/initialize"

// 项目配置
func Setup() {
	// init config
	initialize.Viper()

	// init logger
	initialize.Zap()

	// init datasource
	initialize.Gorm()
}
