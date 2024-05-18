package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	APIKey  string
	BaseURL string
	DB *sql.DB

)

func Init() {
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

	initDB()
}





func initDB() {
	// Retrieve database credentials from environment variables
	dbUser := "root"
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DATABASE")

	// Construct the DSN using the environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Check if the connection is successful.
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
}
