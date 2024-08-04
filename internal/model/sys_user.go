package model

import (
	"errors"
	"time"

	"github.com/xbmlz/starter-gin/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Gender    int       `json:"gender"`
	Phone     string    `json:"phone"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func UpdateUser(user *User) error {
	return db.Get().Save(user).Error
}
