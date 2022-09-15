package request

// User register schemas
type Register struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}
