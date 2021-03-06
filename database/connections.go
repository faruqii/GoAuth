package database

import (
	"log"

	"github.com/faruqii/GoAuth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	connections, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = connections

	err = connections.AutoMigrate(&models.User{}, &models.UserToken{}, &models.Product{}, &models.Merchant{})
	if err != nil {
		return
	}
}
