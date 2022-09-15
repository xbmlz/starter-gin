package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ERROR   = -1
	SUCCESS = 0
)

func Result(c *gin.Context, code int, msg string, data interface{}) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(c, SUCCESS, "操作成功", map[string]interface{}{})
}

func OkWithMsg(c *gin.Context, msg string) {
	Result(c, SUCCESS, msg, map[string]interface{}{})
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, "查询成功", data)
}

func OkWithCustom(c *gin.Context, msg string, data interface{}) {
	Result(c, SUCCESS, msg, data)
}

func Error(c *gin.Context) {
	Result(c, ERROR, "操作失败", map[string]interface{}{})
}

func ErrorWithMsg(c *gin.Context, msg string) {
	Result(c, ERROR, msg, map[string]interface{}{})
}

func ErrorWithCustom(c *gin.Context, data interface{}, msg string) {
	Result(c, ERROR, msg, data)
}
