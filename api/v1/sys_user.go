package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/model"
	"github.com/xbmlz/starter-gin/model/request"
	"github.com/xbmlz/starter-gin/model/response"
	"github.com/xbmlz/starter-gin/service"
	"github.com/xbmlz/starter-gin/utils"
	"go.uber.org/zap"
)

// @Tags SysUser
// @Summary 用户注册
// @Produce  application/json
// @Param data body request.UserRegisterRequest true "用户名, 昵称, 密码, 头像"
// @Success 200 {object} response.Response{data=model.SysUser,msg=string} "用户注册账号,返回包括用户信息"
// @Router /register [post]
func UserRegister(c *gin.Context) {
	var r request.UserRegisterRequest
	_ = c.ShouldBindJSON(&r)
	user := &model.SysUser{
		Username: r.Username,
		NickName: r.NickName,
		Password: r.Password,
		Avatar:   r.Avatar,
	}
	userRet, err := service.UserRegister(*user)
	if err != nil {
		global.Log.Error("注册失败!", zap.Error(err))
		response.ErrorWithMsg(c, "注册失败")
	} else {
		response.OkWithCustom(c, "注册成功", userRet)
	}
}

// @Tags SysUser
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.UserLoginRequest true "用户名, 密码"
// @Success 200 {object} response.Response{data=response.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /login [post]
func UserLogin(c *gin.Context) {
	var r request.UserLoginRequest
	_ = c.ShouldBindJSON(&r)

	user, err := service.UserByUsername(r.Username)
	if err != nil {
		response.ErrorWithMsg(c, "用户名不存在")
	} else {
		if ok := utils.BcryptCheck(r.Password, user.Password); !ok {
			response.ErrorWithMsg(c, "密码错误")
		} else {
			j := utils.NewJWT()
			token, err := j.CreateToken(model.BaseClaims{
				ID:       user.ID,
				Username: user.Username,
				NickName: user.NickName,
			})
			if err != nil {
				global.Log.Error("获取token失败!", zap.Error(err))
				response.ErrorWithMsg(c, "密码错误")
				return
			}
			response.OkWithData(c, response.LoginResponse{
				User:  user,
				Token: token,
			})
		}
	}
}
