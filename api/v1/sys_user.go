package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/model"
	"github.com/xbmlz/starter-gin/model/request"
	"github.com/xbmlz/starter-gin/model/response"
	"github.com/xbmlz/starter-gin/service"
	"go.uber.org/zap"
)

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "用户名, 昵称, 密码, 头像"
// @Success 200 {object} response.Response{data=model.SysUser,msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var r request.Register
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
