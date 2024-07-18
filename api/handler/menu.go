package handler

import (
	"github.com/gin-gonic/gin"
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
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := menuService.CreateMenu(&menu)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, menu)
}

// @Tags Menu
// @Summary List all menus
// @Produce  json
// @Success 200 {array} model.Menu
func GetMenus(c *gin.Context) {
	menus, err := menuService.GetMenus()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, menus)
}
