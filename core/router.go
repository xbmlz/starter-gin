package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	PublicGroup := engine.Group("")
	{
		// 健康检测
		PublicGroup.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
	}
}
