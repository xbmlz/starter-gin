package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xbmlz/starter-gin/docs"
	"github.com/xbmlz/starter-gin/middleware"
)

func RegisterRouter(engine *gin.Engine) {
	// docs
	docs.SwaggerInfo.BasePath = "/api/v1"
	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// middleware
	engine.Use(gin.Recovery())
	// cors 跨域
	engine.Use(middleware.Cors())
	// public
	PublicGroup := engine.Group("")
	{
		// 健康检测
		PublicGroup.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
	}
}
