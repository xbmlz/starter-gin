package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/conf"
	"github.com/xbmlz/starter-gin/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		fields := []zapcore.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if len(c.Errors) > 0 {
			log.Logger.Error(conf.Log.Level, fields...)
		} else {
			log.Logger.Info(conf.Log.Level, fields...)
		}
	}
}
