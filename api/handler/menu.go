package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/starter-gin/api/handler/request"
	"github.com/xbmlz/starter-gin/api/handler/response"
	"github.com/xbmlz/starter-gin/api/model"
)

// @Tags Menu
// @Summary Create a menu
// @Accept  json
// @Produce  json
// @Param menu body model.Menu true "Menu object"
// @Success 200 {object} model.Menu
func CreateMenu(c *gin.Context) {
	var menu model.Menu
	if request.BindJSON(c, &menu) {
		return
	}

	err := menuService.Create(&menu)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Ok(c, menu)
}

// @Tags Menu
// @Summary Update a menu
// @Accept  json
// @Produce  json
// @Param menu body model.Menu true "Menu object"
// @Success 200 {object} model.Menu
func UpdateMenu(c *gin.Context) {
	var menu model.Menu
	if request.BindJSON(c, &menu) {
		return
	}
	err := menuService.Update(&menu)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Ok(c, menu)
}

// @Tags Menu
// @Summary List all menus
// @Produce  json
// @Success 200 {array} model.Menu
func GetMenus(c *gin.Context) {
	menus, err := menuService.List()
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Ok(c, menus)
}
