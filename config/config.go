package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errore durante il caricamento del file .env:", err)
	}

	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value
}
