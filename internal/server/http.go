package server

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/api"
	"github.com/xbmlz/starter-gin/internal/config"
	"github.com/xbmlz/starter-gin/internal/data"
	"go.uber.org/zap"
)

func NewHTTPServer(c *config.Server, logger *zap.Logger, d *data.Data) *http.Server {
	if c.HTTP.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.DateTime, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	apiRouter := r.Group("api")

	api.RegisterHelloRoutes(apiRouter)

	srv := &http.Server{
		Addr:    c.HTTP.Addr(),
		Handler: r,
	}

	return srv
}
