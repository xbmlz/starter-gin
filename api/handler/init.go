package handler

import "github.com/xbmlz/starter-gin/api/service"

var (
	userService = new(service.UserService)
	menuService = new(service.MenuService)
)
