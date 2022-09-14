package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// health check
	r.GET("/ping", api.Ping)
	return r
}
