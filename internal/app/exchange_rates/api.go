package exchange_rates

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/models"
)

// FetchExchangeRates makes an HTTP GET request to the provided API URL and returns the exchange rates.
func FetchExchangeRates(apiURL string) (*models.ExchangeRatesResponse, error) {
    // Make the HTTP GET request to the API
    resp, err := http.Get(apiURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Unmarshal the JSON response into the ExchangeRatesResponse struct
    var exchangeRates models.ExchangeRatesResponse
    err = json.Unmarshal(body, &exchangeRates)
    if err != nil {
        return nil, err
    }

    return &exchangeRates, nil
}
