package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		// Load environment variables if not already loaded
		err := godotenv.Load()
		if err != nil {
			// Handle error if unable to load environment variables
			log.Fatal("Error loading .env file")
		}
		value = os.Getenv(key)
	}
	return value
}
