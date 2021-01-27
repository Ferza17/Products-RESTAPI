package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvironmentVariable(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
