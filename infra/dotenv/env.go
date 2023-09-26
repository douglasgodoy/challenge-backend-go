package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
