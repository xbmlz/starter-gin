package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Body{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Error(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, Body{
		Code:    1,
		Message: message,
		Data:    nil,
	})
}

func Unauthorized(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, Body{
		Code:    401,
		Message: message,
		Data:    nil,
	})
}

func ErrorWithCode(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, Body{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
