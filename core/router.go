package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xbmlz/starter-gin/docs"
)

func RegisterRouter(engine *gin.Engine) {
	// docs
	docs.SwaggerInfo.BasePath = "/api/v1"
	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// public
	PublicGroup := engine.Group("")
	{
		// 健康检测
		PublicGroup.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
	}
}
