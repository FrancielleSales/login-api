package main

import (
	"log"
	"login-api/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	port := config.LoadPort()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run the server: ", err)
	}
}
