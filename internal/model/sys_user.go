package model

import (
	"errors"

	"github.com/xbmlz/starter-gin/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func (u *User) IsAdmin() bool {
	return u.Username == "admin"
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (user *User, err error) {
	err = db.Get().Where("username = ?", username).First(&user).Error
	return user, err
}

func CreateUser(user *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return db.Get().Create(user).Error
}

func VerifyUser(username, password string) (user *User, err error) {
	user, err = GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(password); err != nil {
		return nil, errors.New("incorrect password")
	}
	return user, nil
}

func GetUserByID(id uint) (user *User, err error) {
	err = db.Get().Where("id = ?", id).First(&user).Error
	return user, err
}
