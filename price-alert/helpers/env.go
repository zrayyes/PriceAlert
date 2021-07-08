package helpers

import "os"

// Get environment variable or default to a string
func GetEnv(name string, defaultValue string) string {
	if port, ok := os.LookupEnv(name); ok {
		return port
	}
	return defaultValue
}
