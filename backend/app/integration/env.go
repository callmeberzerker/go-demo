package integration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvValue - gets env value based on a key
func GetEnvValue(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
