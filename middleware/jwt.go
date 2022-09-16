package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/model/response"
	"github.com/xbmlz/starter-gin/utils"
)

// jwt middleware
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtConfig := global.Config.JWT
		token := c.Request.Header.Get(jwtConfig.Header)

		// token是否存在
		if token == "" {
			response.Result(c, http.StatusUnauthorized, "未登录或非法访问", nil)
			c.Abort()
			return
		}

		// token是否合法
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			response.Result(c, http.StatusUnauthorized, err.Error(), nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
