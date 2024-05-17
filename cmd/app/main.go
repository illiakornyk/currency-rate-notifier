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

func main() {

	    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

 	  // Retrieve the API key from the environment variable
    apiKey := os.Getenv("EXCHANGERATESAPI_KEY")

     // Construct the base URL
    baseURL := "http://api.exchangeratesapi.io/v1/latest"

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
    exchangeRatesData, err := exchange_rates.FetchExchangeRates(apiURL)
    if err != nil {
        log.Fatal(err)
    }

	    // Use the exchangeRates data as needed
    // For example, convert EUR to USD and UAH to USD
    usdToUahRate := utils.ConvertEURtoUSDUAH(exchangeRatesData.Rates["USD"], exchangeRatesData.Rates["UAH"])

    // Output the rate for verification
    log.Printf("USD to UAH rate: %f", usdToUahRate)

    // Construct the DSN from the environment variables or directly.
	dsn := "user:password@tcp(localhost:3306)/mydb"

	// Open a connection to the database.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
}
