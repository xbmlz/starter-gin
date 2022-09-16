package service

import (
	"errors"

	"github.com/xbmlz/starter-gin/global"
	"github.com/xbmlz/starter-gin/model"
	"github.com/xbmlz/starter-gin/utils"
	"gorm.io/gorm"
)

func UserRegister(u model.SysUser) (ret model.SysUser, err error) {
	var user model.SysUser
	if !errors.Is(global.DB.First(&user, "username = ?", u.Username).Error, gorm.ErrRecordNotFound) {
		return ret, errors.New("该用户名已被注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	err = global.DB.Create(&u).Error
	return u, err
}

func UserByUsername(username string) (ret model.SysUser, err error) {
	var user model.SysUser
	if errors.Is(global.DB.First(&user, "username = ?", username).Error, gorm.ErrRecordNotFound) {
		return ret, errors.New("用户名不存在")
	}
	return user, err
}
