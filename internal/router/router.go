package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/handlers"
)

func Init(r *gin.Engine) {

	handlers.NewViewHandler().Register(r)

	api := r.Group("/api")
	{
		handlers.NewDemoHandler().Register(api)
	}
}
