package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/exchange_rates"
)

// RatesHandler handles requests for the /rates route
func RatesHandler(w http.ResponseWriter, r *http.Request) {
	// Create a URL object from the base URL
	u, err := url.Parse(config.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare query parameters
	q := u.Query()
	q.Set("access_key", config.APIKey)
	q.Set("symbols", "USD,UAH")
	q.Set("format", "1")

	// Assign encoded query parameters to the URL object
	u.RawQuery = q.Encode()

	// The fully constructed URL with the embedded API key and other parameters
	apiURL := u.String()

	// Fetch the exchange rates data
	exchangeRatesData, err := exchange_rates.FetchExchangeRates(apiURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract the USD to UAH rate
	usdToUahRate := exchangeRatesData.Rates["UAH"]

	// Create a struct to format the response
	type response struct {
		Rate float64 `json:"rate"`
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the response as JSON
	json.NewEncoder(w).Encode(response{Rate: usdToUahRate})
}
