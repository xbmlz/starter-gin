package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/model"
)

type userHandler struct {
	BaseHandler
}

type UpdateUserRequest struct {
	Nickname  string `json:"nickname"`
	Gender    int    `json:"gender"`
	AvatarURL string `json:"avatar_url"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func NewUserHandler() userHandler {
	return userHandler{}
}

func (h userHandler) Register(router *gin.RouterGroup) {
	router.GET("/user", h.GetUser)
	router.PATCH("/user", h.UpdateUser)
}

func (h userHandler) GetUser(c *gin.Context) {
	h.Ok(c, h.GetCurrentUser(c))
}

func (h userHandler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	if h.BindJSON(c, &req) {
		h.BadRequest(c, "Invalid request body")
		return
	}

	user := h.GetCurrentUser(c)
	user.Nickname = req.Nickname
	user.Gender = req.Gender
	user.AvatarURL = req.AvatarURL
	user.Phone = req.Phone
	user.Email = req.Email

	err := model.UpdateUser(user)
	if err != nil {
		h.Error(c, err.Error())
		return
	}

	h.Ok(c, user)
}
