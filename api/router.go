package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/api/handler"
)

func RegiesterRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// public routes
	api := router.Group("/api")
	{
		api.POST("/login", handler.Login)
		api.POST("/register", handler.Register)
		api.POST("/logout", handler.Logout)
	}
}
