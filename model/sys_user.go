package model

type SysUser struct {
	ID       uint   `json:"id" gorm:"primaryKey;auto_increment;comment:id"`
	Username string `json:"userName" gorm:"comment:用户登录名"`
	Password string `json:"-"  gorm:"comment:用户登录密码"`
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Avatar   string `json:"avatar" gorm:"comment:用户头像"`
	Email    string `json:"email"  gorm:"comment:用户邮箱"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
