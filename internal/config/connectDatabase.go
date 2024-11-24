package config

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global variable that stores the database connection
var DB *gorm.DB

// Function to connect to the database
func ConnectDatabase() {
	host, user, password, dbname, port := LoadDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	DB = db
	fmt.Println("Connected to the database!")
}
