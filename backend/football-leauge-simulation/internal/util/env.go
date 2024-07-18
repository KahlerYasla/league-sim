package util

import (
	"os"
	"strconv"
)

// GetEnv retrieves the value of the environment variable named by the key.
// It returns the value, which will be either the value of the environment variable or the fallback value if the variable is not present.
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetEnvAsInt retrieves the value of the environment variable named by the key as an integer.
// It returns the value, which will be either the value of the environment variable or the fallback value if the variable is not present or cannot be converted to an integer.
func GetEnvAsInt(name string, fallback int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}

// GetEnvAsBool retrieves the value of the environment variable named by the key as a boolean.
// It returns the value, which will be either the value of the environment variable or the fallback value if the variable is not present or cannot be converted to a boolean.
func GetEnvAsBool(name string, fallback bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return fallback
}
