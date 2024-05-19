package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	APIKey  string
	BaseURL string
	DB *sql.DB

)

func Init() {
	// Load the .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// Retrieve the API key from the environment variable
	APIKey = os.Getenv("EXCHANGERATESAPI_KEY")
	if APIKey == "" {
		log.Fatal("EXCHANGERATESAPI_KEY is not set")
	}

	// Retrieve the base URL from the environment variable
	BaseURL = os.Getenv("EXCHANGERATESAPI_BASE_URL")
	if BaseURL == "" {
		log.Fatal("EXCHANGERATESAPI_BASE_URL is not set")
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

	// Attempt to connect to the database with retries
	for i := 0; i < 5; i++ {
		DB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Error opening database: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Check if the connection is successful.
		err = DB.Ping()
		if err != nil {
			log.Printf("Error connecting to the database: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("Successfully connected to the database!")
		return
	}

	log.Fatal("Unable to connect to the database after several retries.")
}
