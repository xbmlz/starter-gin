package response

import "github.com/xbmlz/starter-gin/model"

type LoginResponse struct {
	User  model.SysUser `json:"user"`
	Token string        `json:"token"`
}
