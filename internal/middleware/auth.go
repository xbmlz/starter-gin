package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/constant"
	"github.com/xbmlz/starter-gin/pkg/auth"
)

func JWTAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractToken(c)
		if len(tokenString) == 0 {
			AbortHTML(c)
		}
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			AbortHTML(c)
		}
		c.Set(constant.CurrentUserKey, claims)
		c.Next()
	}
}

func SessionAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get(constant.SessionUserKey) == nil {
			AbortHTML(c)
		}

		c.Set(constant.CurrentUserKey, session.Get(constant.SessionUserKey))
		c.Next()
	}
}

func ExtractToken(ctx *gin.Context) (token string) {
	token = ctx.GetHeader(constant.TokenName)
	if len(token) == 0 {
		token = ctx.Query(constant.TokenName)
	}
	return strings.TrimPrefix(token, "Bearer ")
}

func AbortHTML(c *gin.Context) {
	// redirect to login page
	c.Redirect(http.StatusFound, "/login.html")
	c.Abort()
}

func AbortJSON(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": http.StatusUnauthorized,
		"msg":  "Unauthorized",
		"data": nil,
	})
	c.Abort()
}
