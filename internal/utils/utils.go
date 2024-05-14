package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Basic "contains string" function as a placeholder - this will be filled in later
func ContainsString(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

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
