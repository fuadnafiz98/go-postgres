package database

import (
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

	// database.AutoMigrate(&models.Message{})

	DB = database
}
