package controllers

import (
	"net/http"

	"github.com/Arghya-Banerjee/gocrud-api/database"
	"github.com/Arghya-Banerjee/gocrud-api/models"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}