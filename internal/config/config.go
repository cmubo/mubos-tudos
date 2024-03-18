package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Error(err)
	}

	return os.Getenv(key)
}
