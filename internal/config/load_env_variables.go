package config

import (
	"log"

	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables")
	}
}
