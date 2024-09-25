package controllers

import (
	"net/http"

	"github.com/Arghya-Banerjee/gocrud-api/database"
	"github.com/Arghya-Banerjee/gocrud-api/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}