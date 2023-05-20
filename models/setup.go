package models

import (
	"github.com/farwater-create/backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DSNOptions struct {
	MARIADB_USER     string
	MARIADB_PASSWORD string
	MARIADB_HOST     string
}

func init() {
	ConnectDatabase()
}

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(config.Environment.MARIADB_DSN))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&ApiKey{})

	DB = database
}

func AssertInputType[T any](input any) T {
	i, ok := input.(T)
	if !ok {
		panic("invalid input type error, did you use the wrong Input type for a model?")
	}
	return i
}
