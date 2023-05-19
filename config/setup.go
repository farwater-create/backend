package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	MARIADB_DSN string `validate:"required"`
	PORT        string `validate:"required"`
}

var Environment Config = Config{
	MARIADB_DSN: os.Getenv("MARIADB_DSN"),
	PORT:        os.Getenv("PORT"),
}

func init() {
	godotenv.Load()
	validate := validator.New()
	err := validate.Struct(Environment)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}
