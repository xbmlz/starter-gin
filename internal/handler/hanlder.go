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

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}

func (e Error) Error() string {
	return e.Message
}

func HandleResponse(ctx *gin.Context, err error, data interface{}) {
	if err == nil {
		ctx.JSON(http.StatusOK, Response{
			Code:    http.StatusOK,
			Message: "success",
			Data:    data,
		})
		return
	}

	if code, ok := errorCodeMap[err]; ok {
		ctx.JSON(code, Response{
			Code:    code,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		Data:    nil,
	})
}

func BindAndCheck(ctx *gin.Context, req interface{}) bool {
	if err := ctx.ShouldBind(req); err != nil {
		log.Sugar.Error("http_handle BindAndCheck fail, %s", err.Error())
		HandleResponse(ctx, errors.New("invalid request"), nil)
		return true
	}
	return false
}
