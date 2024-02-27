package model

import (
	"errors"

	"github.com/glebarez/sqlite"
	"github.com/xbmlz/starter-gin/internal/conf"
	"github.com/xbmlz/starter-gin/internal/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() error {

	var (
		db  *gorm.DB
		err error
	)

	switch conf.Database.Type {
	case "postgres":
		db, err = gorm.Open(postgres.Open(conf.Database.DSN), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(conf.Database.DSN), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(conf.Database.DSN), &gorm.Config{})
	default:
		return errors.New("unsupported database type")
	}

	if err != nil {
		return err
	}

	DB = db.Debug()

	if err := DB.AutoMigrate(&User{}); err != nil {
		return err
	}

	log.Sugar.Info("database auto migrate success")

	return nil
}
