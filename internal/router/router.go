package router

import (
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/constant"
	"github.com/xbmlz/starter-gin/internal/db"
	"github.com/xbmlz/starter-gin/internal/handlers"
	"github.com/xbmlz/starter-gin/internal/middleware"
)

func Init(r *gin.Engine) {

	store := gormsessions.NewStore(db.Get(), true, []byte(constant.SessionKey))
	r.Use(sessions.Sessions(constant.SessionName, store))

	handlers.NewViewHandler().Register(r)
	handlers.NewAuthHandler().Register(r)

	authRouter := r.Group("/api", middleware.JWTAuthRequired())
	{
		handlers.NewDemoHandler().Register(authRouter)
	}
}
