package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
// @Success 200 {object} LoginResponse
func Login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := userService.VerifyUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := service.TokenGenerate(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: fmt.Sprintf("Bearer %s", token)})
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
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIn := &model.User{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := userService.CreateUser(userIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func Logout(c *gin.Context) {

}
