package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func Config(key string) string {
	// TODO: there is a problem with this in tests, since any tests create a temporary directory
	// When a test is ran from a directory, it cant find the env file because it isnt there.
	err := godotenv.Load()

	if err != nil {
		log.Error(err)
	}

	return os.Getenv(key)
}
