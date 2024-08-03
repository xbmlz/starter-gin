package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/constant"
	"github.com/xbmlz/starter-gin/internal/model"
)

type authHandler struct {
	BaseHandler
}

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func NewAuthHandler() authHandler {
	return authHandler{}
}

func (h authHandler) Register(router *gin.Engine) {
	router.POST("/login", h.login)
	router.POST("/logout", h.logout)
}

func (h authHandler) login(c *gin.Context) {
	var req LoginRequest
	if h.Bind(c, &req) {
		h.BadRequest(c, "Invalid request")
		return
	}

	user, err := model.VerifyUser(req.Username, req.Password)
	if err != nil {
		h.Unauthorized(c, "Invalid username or password")
		return
	}

	// Generate JWT token
	// token, err := auth.GenerateToken(user.ID, user.Username)
	// if err != nil {
	// 	h.Error(c, err.Error())
	// 	return
	// }

	// create session
	session := sessions.Default(c)
	session.Set(constant.SessionUserKey, user.ID)
	session.Save()

	h.Ok(c, nil)
}

func (h authHandler) logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(constant.SessionUserKey)
	session.Save()
	h.Ok(c, nil)
}
