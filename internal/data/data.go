package data

import (
	"errors"

	"github.com/glebarez/sqlite"
	"github.com/redis/go-redis/v9"
	"github.com/xbmlz/starter-gin/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewData(c *config.Data, logger *zap.Logger) (*Data, func(), error) {

	// initialize database connection
	var dialector gorm.Dialector
	switch c.Database.Driver {
	case "sqlite3":
		dialector = sqlite.Open(c.Database.DSN)
	case "mysql":
		dialector = mysql.Open(c.Database.DSN)
	case "postgres":
		dialector = postgres.Open(c.Database.DSN)
	default:
		return nil, nil, errors.New("unsupported database dialect")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logger.Sugar().Errorf("failed to connect to database: %v", err)
		return nil, nil, err
	}

	// initialize redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Password,
		DB:       int(c.Redis.DB),
	})

	d := &Data{
		db:  db,
		rdb: rdb,
	}

	return d, func() {
		if err := d.rdb.Close(); err != nil {
			logger.Sugar().Errorf("failed to close redis connection: %v", err)
		}
	}, nil

}
