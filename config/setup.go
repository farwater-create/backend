package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type Config struct {
	MARIADB_DSN             string `validate:"required"`
	PORT                    string `validate:"required"`
	KOFI_VERIFICATION_TOKEN string `validate:"required"`
}

var Environment Config = Config{
	MARIADB_DSN:             os.Getenv("MARIADB_DSN"),
	PORT:                    os.Getenv("PORT"),
	KOFI_VERIFICATION_TOKEN: os.Getenv("KOFI_VERIFICATION_TOKEN"),
}

func init() {
	godotenv.Load()
	validate := validator.New()
	err := validate.Struct(Environment)
	logrus.Error(err)
	if err != nil {
		panic(err)
	}
}
