package utils

import (
	"log"
	"os"
	"strings"
)

var secretPaths map[string]string = map[string]string{
	"KAFKA_BROKERS":        os.Getenv("KAFKA_BROKERS_FILE"),
	"TWITCH_CLIENT_ID":     os.Getenv("TWITCH_CLIENT_ID_FILE"),
	"TWITCH_CLIENT_SECRET": os.Getenv("TWITCH_CLIENT_SECRET_FILE"),
}

var secrets map[string]string

func ReadSecrets() map[string]string {
	for secretName, secretPath := range secretPaths {
		if secrets[secretName] == "" {
			secret, err := os.ReadFile(secretPath)
			if err != nil {
				log.Fatalf("Failed to read secret from %s: %v", secretPath, err)
			}
			secrets[secretName] = strings.TrimSpace(string(secret))
		}
	}

	return secrets
}
