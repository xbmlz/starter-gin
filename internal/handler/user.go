package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterReq struct {
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

// Register godoc
// @Summary 用户注册
// @Schemes
// @Description 目前只支持邮箱登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body UserRegisterReq true "params"
// @Success 200 {object} Response
// @Router /register [post]
func Register(ctx *gin.Context) {
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
