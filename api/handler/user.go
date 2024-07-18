package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags Users
// @Summary Get User
// @Description Get User by ID
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
func GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.Atoi(id)

	user, err := userService.GetUserByID(userId)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}
