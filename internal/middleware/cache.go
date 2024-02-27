package middleware

import "github.com/gin-gonic/gin"

func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/" {
			c.Header("Cache-Control", "no-cache")
		} else {
			c.Header("Cache-Control", "max-age=604800")
		}
		c.Next()
	}
}
