package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/exchange_rates"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/utils"
)

// ConstructAPIURL constructs the URL to fetch exchange rates.
func ConstructAPIURL() (string, error) {
	u, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("access_key", config.APIKey)
	q.Set("symbols", "USD,UAH")
	q.Set("format", "1")

	u.RawQuery = q.Encode()

	return u.String(), nil
}

// FetchAndRespondExchangeRate fetches the exchange rates and writes the response.
func FetchAndRespondExchangeRate(w http.ResponseWriter, apiURL string) {
	exchangeRatesData, err := exchange_rates.FetchExchangeRates(apiURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usdToUahRate := utils.ConvertEURtoUSDUAH(exchangeRatesData.Rates["USD"], exchangeRatesData.Rates["UAH"])

	type response struct {
		Rate float64 `json:"rate"`
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Rate: usdToUahRate})
}

// RatesHandler handles requests for the /rates route.
func RatesHandler(w http.ResponseWriter, r *http.Request) {
	apiURL, err := ConstructAPIURL()
	if err != nil {
		log.Fatal(err)
	}

	FetchAndRespondExchangeRate(w, apiURL)
}
