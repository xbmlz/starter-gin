package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/api/handler/response"
)

func BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		response.ErrorWithCode(c, http.StatusBadRequest, err.Error())
		return true
	}
	return false
}
