package config

import (
	"os"
)

func Config(key string) string {
	return os.Getenv(key)
}

func GetEnv(key string) string {

	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value
}
