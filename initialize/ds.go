package initialize

import (
	"os"

	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 初始化数据库 https://github.com/go-gorm/gorm
func InitDatasource() {
	var db *gorm.DB
	switch global.Config.Datasource.Primary {
	case "mysql":
		db = DsMysql()
	default:
		db = DsSqlite()
	}
	global.DB = db
	if global.DB != nil {
		autoMigrate(global.DB)
	}
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
	)
	if err != nil {
		global.Log.Error("auto migrate table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("auto migrate table success")
}
