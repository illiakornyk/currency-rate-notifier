package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	APIKey  string
	BaseURL string
)

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the API key from the environment variable
	APIKey = os.Getenv("EXCHANGERATESAPI_KEY")
	if APIKey == "" {
		log.Fatal("EXCHANGERATESAPI_KEY is not set in .env file")
	}

	// Retrieve the base URL from the environment variable
	BaseURL = os.Getenv("EXCHANGERATESAPI_BASE_URL")
	if BaseURL == "" {
		log.Fatal("EXCHANGERATESAPI_BASE_URL is not set in .env file")
	}
}
