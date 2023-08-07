package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Helps retrieve the MongoDB URI fron the env file
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGOURI")
}
