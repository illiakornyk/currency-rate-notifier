package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/exchange_rates"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/utils"
	"github.com/joho/godotenv"
)

var (
	apiKey string
	baseURL string
	db *sql.DB
)

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the API key from the environment variable
	apiKey = os.Getenv("EXCHANGERATESAPI_KEY")
	if apiKey == "" {
		log.Fatal("EXCHANGERATESAPI_KEY is not set in .env file")
	}

	// Retrieve the base URL from the environment variable
	baseURL = os.Getenv("EXCHANGERATESAPI_BASE_URL")
	if baseURL == "" {
		log.Fatal("EXCHANGERATESAPI_BASE_URL is not set in .env file")
	}

	// Initialize the database connection
	initDB()
}

func initDB() {
	// Construct the DSN from the environment variables or directly.
	dsn := "user:password@tcp(localhost:3306)/mydb"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Check if the connection is successful.
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
}

func main() {
    // Create a URL object from the base URL
    u, err := url.Parse(baseURL)
    if err != nil {
        log.Fatal(err)
    }

    // Prepare query parameters
    q := u.Query()
    q.Set("access_key", apiKey)
    q.Set("symbols", "USD,UAH")
    q.Set("format", "1")

    // Assign encoded query parameters to the URL object
    u.RawQuery = q.Encode()

    // The fully constructed URL with the embedded API key and other parameters
    apiURL := u.String()

    // Fetch the exchange rates data
    exchangeRatesData, err := exchange_rates.FetchExchangeRates(apiURL)
    if err != nil {
        log.Fatal(err)
    }

    // Convert EUR to USD and UAH to USD
    usdToUahRate := utils.ConvertEURtoUSDUAH(exchangeRatesData.Rates["USD"], exchangeRatesData.Rates["UAH"])

    // Output the rate for verification
    log.Printf("USD to UAH rate: %f", usdToUahRate)


    fmt.Println("Operation completed successfully.")
}
