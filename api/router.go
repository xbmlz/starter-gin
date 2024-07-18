package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xbmlz/starter-gin/api/handler"
	"github.com/xbmlz/starter-gin/api/middlerware"
)

func RegiesterRouter(router *gin.Engine) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// public routes
	api := router.Group("/api")

	noAuth := api.Group("")
	{
		noAuth.POST("/login", handler.Login)
		noAuth.POST("/register", handler.Register)
		noAuth.POST("/logout", handler.Logout)
	}

	auth := api.Group("").Use(middlerware.Auth())
	{
		auth.GET("/users/:id", handler.GetUser)

		auth.POST("/menus", handler.CreateMenu)
		auth.GET("/menus", handler.GetMenus)
	}
}
