package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/handler"
	"github.com/xbmlz/starter-gin/internal/log"
	"github.com/xbmlz/starter-gin/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			log.Sugar.Warn("token is empty, url")
			handler.HandleResponse(ctx, handler.ErrUnauthorized, nil)
			ctx.Abort()
			return
		}

		_, err := jwt.ParseToken(tokenString)
		if err != nil {
			log.Sugar.Warn("parse token failed")
			handler.HandleResponse(ctx, handler.ErrUnauthorized, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
