package controllers

import (
	"net/http"

	"github.com/Arghya-Banerjee/gocrud-api/database"
	"github.com/Arghya-Banerjee/gocrud-api/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
