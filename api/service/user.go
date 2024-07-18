package service

import (
	"errors"

	"github.com/xbmlz/starter-gin/api/model"
	"github.com/xbmlz/starter-gin/internal/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	// find user by username
	var u model.User
	if err := db.Get().Where("username = ?", user.Username).First(&u).Error; err != gorm.ErrRecordNotFound {
		return nil, errors.New("username already exists")
	}

	// create user
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashPwd)
	if err := db.Get().Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) VerifyUser(username, password string) (user *model.User, err error) {
	// find user by username
	err = db.Get().Where("username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

func (s *UserService) GetUserByID(id int) (*model.User, error) {
	var user model.User
	if err := db.Get().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
