package model

type User struct {
	Id       int64 `gorm:"primaryKey;autoIncrement"`
	Nickname string
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Insert() error {
	return DB.Create(u).Error
}

func GetUserByEmail(email string) (user *User, err error) {
	if err = DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers() (users []User, err error) {
	if err = DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
