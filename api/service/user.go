package service

import (
	"errors"

	"github.com/xbmlz/starter-gin/api/model"
	"github.com/xbmlz/starter-gin/internal/db"
	"gorm.io/gorm"
)

func UserCreate(user *model.User) (*model.User, error) {
	// find user by username
	var u model.User
	if err := db.Get().Where("username =?", user.Username).First(&u).Error; err != gorm.ErrRecordNotFound {
		return nil, errors.New("username already exists")
	}

	// create user
	user.HashPassword()
	if err := db.Get().Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
