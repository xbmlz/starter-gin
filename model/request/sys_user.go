package request

// User register schemas
type UserRegisterRequest struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

// User register schemas
type UserLoginRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
