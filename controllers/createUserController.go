package controllers

import (
	"net/http"
	"log"

	"github.com/Arghya-Banerjee/gocrud-api/database"
	"github.com/Arghya-Banerjee/gocrud-api/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	if database.DB == nil {
        log.Println("Database connection is not initialized")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
        return
    }

    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, user)
}
