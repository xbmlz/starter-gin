package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/xbmlz/starter-gin/api/v1"
	"github.com/xbmlz/starter-gin/docs"
	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/middleware"
)

func RegisterRouter() *gin.Engine {
	// 设置服务模式
	gin.SetMode(global.Config.Server.Mode)
	engine := gin.New()
	// docs
	docs.SwaggerInfo.BasePath = "/"
	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// middleware
	engine.Use(gin.Recovery())
	// cors 跨域
	engine.Use(middleware.Cors())
	// root
	rootGroup := engine.Group("")
	{
		// 健康检测
		rootGroup.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
		// 注册
		rootGroup.POST("/register", v1.UserRegister)
		// 登录
		rootGroup.POST("/login", v1.UserLogin)
	}
	// v1
	// v1Group := engine.Group("/api/v1").Use(middleware.JWTAuth())
	// {
	// 	// user

	// }
	return engine
}
