package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/constant"
	"github.com/xbmlz/starter-gin/internal/model"
)

type BaseHandler struct {
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (h *BaseHandler) Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func (h *BaseHandler) Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, response{
		Code: 1,
		Msg:  msg,
		Data: nil,
	})
}

func (h *BaseHandler) ErrorWithCode(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func (h *BaseHandler) Unauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, response{
		Code: 401,
		Msg:  msg,
		Data: nil,
	})
}

func (h *BaseHandler) BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, response{
		Code: 400,
		Msg:  msg,
		Data: nil,
	})
}

func (h *BaseHandler) BindJSON(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		return true
	}
	return false
}

func (h *BaseHandler) Bind(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBind(obj); err != nil {
		return true
	}
	return false
}

func (h *BaseHandler) GetCurrentUser(c *gin.Context) *model.User {
	// Get user ID from session
	session := sessions.Default(c)
	userID := session.Get(constant.SessionUserKey)
	if userID == nil {
		h.Unauthorized(c, "Unauthorized")
		return nil
	}

	// Get user from database
	user, err := model.GetUserByID(userID.(uint))
	if err != nil {
		h.Error(c, err.Error())
		return nil
	}

	return user
}
