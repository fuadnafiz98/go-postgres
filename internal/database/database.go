package database

import (
	models "github.com/fuadnafiz98/go-postgres/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=admin password=fuad1234 dbname=go-postgres port=5432 sslmode=disable TimeZone=Asia/Dhaka"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database!")
	}

	database.AutoMigrate(&models.Message{})

	DB = database
}
