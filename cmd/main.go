package main

import (
	"log"
	"login-api/internal/config"
	"login-api/pkg/handlers"
	"login-api/pkg/repositories"
	"login-api/pkg/services"
	"github.com/gin-gonic/gin"
)

// Connect to the database and synchronize migrations
func init() {
	config.ConnectDatabase()
	config.SyncDatabase()
}

// Create users repository, service, and handler instances
func getUserHandler() *handlers.UserHandler {
	db := config.DB
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	return handlers.NewUserHandler(userService)
}

func main() {
	// Create users handler
	usersHandler := getUserHandler()

	// Create the Gin server
	r := gin.Default()

	// Define routes
	v1 := r.Group("/v1")
	{
		// Endpoint for creating users
		v1.POST("/users", usersHandler.CreateUser)
	}

	// Load port from configuration
	port := config.LoadPort()

	// Start server
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run the server: ", err)
	}
}
