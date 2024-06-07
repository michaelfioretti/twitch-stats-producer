package utils

import (
	"os"
	"testing"
)

func TestGetEnvVar_ExistingKey(t *testing.T) {
	key := "EXISTING_KEY"
	value := "existing_value"
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}

	defer os.Unsetenv(key)

	result := GetEnvVar(key)

	if result != value {
		t.Errorf("Expected %s, but got %s", value, result)
	}
}
