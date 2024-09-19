package api

import "github.com/gin-gonic/gin"

func RegisterHelloRoutes(r *gin.RouterGroup) {
	r.GET("/hello", hello)
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
