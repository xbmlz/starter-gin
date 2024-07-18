package middlerware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/api/service"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ExtractToken(c)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, err := service.TokenParse(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func ExtractToken(ctx *gin.Context) (token string) {
	token = ctx.GetHeader("Authorization")
	if len(token) == 0 {
		token = ctx.Query("Authorization")
	}
	return strings.TrimPrefix(token, "Bearer ")
}
