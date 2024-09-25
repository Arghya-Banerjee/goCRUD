package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Arghya-Banerjee/gocrud-api/controllers"
)

// SetupRouter sets up all the application routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.POST("/users", controllers.CreateUser)    	// Create
	r.GET("/users", controllers.GetUsers)       	// Read all
	r.GET("/users/:id", controllers.GetUser)    	// Read one
	r.PUT("/users/:id", controllers.UpdateUser) 	// Update
	r.DELETE("/users/:id", controllers.DeleteUser)  // Delete

	return r
}