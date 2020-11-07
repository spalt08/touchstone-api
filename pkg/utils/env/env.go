// Package env provides a safe access to environment variables with fallback values
package env

import (
	"fmt"
	"os"
)

// GetServingPort returns application port to be served
func GetServingPort() string {
	return getStringValueWithFallback("SERVING_PORT", ":8080")
}

// GetDatabaseAddress returns postgres database port
func GetDatabaseAddress() string {
	return getStringValueWithWarning("POSTGRES_ADDR")
}

// GetDatabaseUser returns postgres database user
func GetDatabaseUser() string {
	return getStringValueWithWarning("POSTGRES_USER")
}

// GetDatabasePassword returns postgres database password
func GetDatabasePassword() string {
	return getStringValueWithWarning("POSTGRES_PASSWORD")
}

// GetDatabaseName returns postgres database name
func GetDatabaseName() string {
	return getStringValueWithWarning("POSTGRES_DB")
}

// Get returns postgres database name
func GetJwtSecret() string {
	return getStringValueWithWarning("MASTER_SECRET")
}

func getStringValueWithFallback(name string, fallback string) string {
	var value = os.Getenv(name)

	if len(value) == 0 {
		value = fallback
	}

	return value
}

func getStringValueWithWarning(name string) string {
	var value = os.Getenv(name)

	if len(value) == 0 {
		panic(fmt.Sprintf("Environment variable is not set: %s", name))
	}

	return value
}
