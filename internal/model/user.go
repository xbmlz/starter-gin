package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

func (u *User) TableName() string {
	return "users"
}
