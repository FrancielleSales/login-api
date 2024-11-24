package config

import (
	"log"
	"login-api/pkg/models"
)

// Synchronizes the tables in the database
func SyncDatabase() {
	err := DB.AutoMigrate(&models.Users{})
	if err != nil {
		log.Fatal("Failed to sync database: ", err)
	}
	log.Println("Database migration completed successfully")
}
