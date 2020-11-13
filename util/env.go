package util

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv will load .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
