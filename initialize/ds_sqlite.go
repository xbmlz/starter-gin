package initialize

import (
	"github.com/xbmlz/starter-gin/global"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DsSqlite() *gorm.DB {
	s := global.Config.Datasource.Sqlite
	if s.Path == "" {
		return nil
	}
	// TODO gorm config
	if db, err := gorm.Open(sqlite.Open(s.Path), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		return db
	}
}
