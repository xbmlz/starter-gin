package handlers

import "github.com/gin-gonic/gin"

type demoHandler struct {
	BaseHandler
}

func NewDemoHandler() demoHandler {
	return demoHandler{}
}

func (d demoHandler) Register(r *gin.RouterGroup) {
	r.GET("/ping", d.ping)
}

func (d demoHandler) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
