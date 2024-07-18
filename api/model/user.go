package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func (u *User) TableName() string {
	return "sys_user"
}
