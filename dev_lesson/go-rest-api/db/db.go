package db

import (
	"fmt"
	"log"

	"github.com/m724tmy/go-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=mypassword dbname=go_rest_api port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected!")
	DB.AutoMigrate(&models.User{})
}