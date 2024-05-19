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

// FetchExchangeRateData fetches the exchange rate data and returns it.
func FetchExchangeRateData(apiURL string) (float64, error) {
	exchangeRatesData, err := exchange_rates.FetchExchangeRates(apiURL)
	if err != nil {
		return 0, err
	}

	usdToUahRate := utils.ConvertEURtoUSDUAH(exchangeRatesData.Rates["USD"], exchangeRatesData.Rates["UAH"])
	return usdToUahRate, nil
}

// RatesHandler handles requests for the /rates route.
func RatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	apiURL, err := ConstructAPIURL()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	usdToUahRate, err := FetchExchangeRateData(apiURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type response struct {
		Rate float64 `json:"rate"`
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Rate: usdToUahRate})
}
