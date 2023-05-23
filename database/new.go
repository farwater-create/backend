package database

import (
	"github.com/farwater-create/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(dsn string) *gorm.DB {
	database, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.ApiKey{})
	return database
}
