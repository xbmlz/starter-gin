package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/api/handler/request"
	"github.com/xbmlz/starter-gin/api/handler/response"
	"github.com/xbmlz/starter-gin/api/model"
	"github.com/xbmlz/starter-gin/api/service"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Tags Auth
// @Summary Login
// @Description Login with username and password
// @Accept  json
// @Produce  json
// @Param body body LoginRequest true "Login Request"
// @Success 200 {object} response.Body{LoginResponse}
func Login(c *gin.Context) {
	req := LoginRequest{}
	if request.BindJSON(c, &req) {
		return
	}
	user, err := userService.VerifyUser(req.Username, req.Password)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	token, err := service.TokenGenerate(user)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Ok(c, LoginResponse{Token: token})
}

// @Tags Auth
// @Summary Register
// @Description Register a new user
// @Accept  json
// @Produce  json
// @Param body body RegisterRequest true "Register Request"
// @Success 200 {object} model.User
func Register(c *gin.Context) {
	req := RegisterRequest{}
	if request.BindJSON(c, &req) {
		return
	}

	userIn := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := userService.CreateUser(userIn)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Ok(c, user)
}

func Logout(c *gin.Context) {

}
