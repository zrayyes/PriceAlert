package helpers

import "os"

func GetEnv(name string, defaultValue string) string {
	if port, ok := os.LookupEnv(name); ok {
		return port
	}
	return defaultValue
}
