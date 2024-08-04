package handlers

import (
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xbmlz/starter-gin/internal/constant"
	"github.com/xbmlz/starter-gin/internal/model"
	"github.com/xbmlz/starter-gin/pkg/env"
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
	router.POST("/upload", h.upload)
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

func (h authHandler) upload(c *gin.Context) {

	uploadDir := env.GetString("UPLOAD_DIR", "uploads")

	// upload file
	file, err := c.FormFile("file")
	if err != nil {
		h.BadRequest(c, "Invalid file")
		return
	}

	ext := filepath.Base(file.Filename)
	filename := uuid.New().String() + "." + ext

	dst := filepath.Join(uploadDir, filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		h.Error(c, err.Error())
		return
	}

	// return file URL TODO: fix this
	h.Ok(c, gin.H{"url": "http://localhost:8080/uploads/" + filename})
}
