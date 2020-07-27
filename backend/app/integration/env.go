package integration

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvValue(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
