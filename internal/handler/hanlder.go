package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/log"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, Response{
			Code:    http.StatusOK,
			Message: "success",
			Data:    data,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

}

func BindAndCheck(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBind(req); err != nil {
		log.Sugar.Error("http_handle BindAndCheck fail, %s", err.Error())
		HandleResponse(ctx, errors.New("invalid request"), nil)
		return true
	}
	return false
}
