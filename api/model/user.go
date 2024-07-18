package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func (u *User) HashPassword() {
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashPwd)
}
