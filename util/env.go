package util

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv will load .env file
func LoadEnv() {
	path := "/home/divierda/Golang Practice/todo-app/.env"
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file, ", err)
	}
}
