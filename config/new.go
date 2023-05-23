package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	MARIADB_DSN             string `validate:"required"`
	PORT                    string `validate:"required"`
	KOFI_VERIFICATION_TOKEN string `validate:"required"`
	SALT                    string `validate:"required"`
}

func New() (Config, error) {
	godotenv.Load()
	config := Config{
		MARIADB_DSN:             os.Getenv("MARIADB_DSN"),
		PORT:                    os.Getenv("PORT"),
		KOFI_VERIFICATION_TOKEN: os.Getenv("KOFI_VERIFICATION_TOKEN"),
		SALT:                    os.Getenv("SALT"),
	}
	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		return config, err
	}
	return config, nil
}
