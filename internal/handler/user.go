package handler

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/conf"
	"github.com/xbmlz/starter-gin/internal/model"
	"github.com/xbmlz/starter-gin/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterReq struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

// UserRegister godoc
// @Summary 用户注册
// @Schemes
// @Description 用户注册
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body UserRegisterReq true "params"
// @Success 200 {object} Response
// @Router /register [post]
func UserRegister(ctx *gin.Context) {
	var req UserRegisterReq
	if BindAndCheck(ctx, &req) {
		return
	}

	if user, err := model.GetUserByEmail(req.Email); err == nil && user != nil {
		HandleResponse(ctx, errors.New("user exist"), nil)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		HandleResponse(ctx, err, nil)
		return
	}

	u := model.User{
		Email:    req.Email,
		Password: string(hashed),
	}

	if err := u.Insert(); err != nil {
		HandleResponse(ctx, err, nil)
		return
	}

	HandleResponse(ctx, nil, nil)
}

// UserLogin godoc
// @Summary 用户登录
// @Schemes
// @Description 用户登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body UserRegisterReq true "params"
// @Success 200 {object} Response
// @Router /login [post]
func UserLogin(ctx *gin.Context) {
	var req UserRegisterReq
	if BindAndCheck(ctx, &req) {
		return
	}

	user, err := model.GetUserByEmail(req.Email)
	if err != nil {
		HandleResponse(ctx, err, nil)
		return
	}

	if user == nil {
		HandleResponse(ctx, errors.New("user not exist"), nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		HandleResponse(ctx, errors.New("password error"), nil)
		return
	}

	// generate token
	exp := time.Now().Add(time.Duration(conf.Config.Security.JWTExpire) * time.Second)
	token, err := jwt.GenToken(strconv.Itoa(int(user.Id)), exp)
	if err != nil {
		HandleResponse(ctx, err, nil)
		return
	}

	HandleResponse(ctx, nil, LoginResponse{AccessToken: token})
}

// GetUsers
// @Summary 用户列表
// @Schemes
// @Description 用户列表
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /users [get]
func GetUsers(ctx *gin.Context) {
	users, err := model.GetUsers()
	if err != nil {
		HandleResponse(ctx, err, nil)
		return
	}
	HandleResponse(ctx, nil, users)
}
