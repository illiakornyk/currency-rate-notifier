package exchange_rates

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/models"
)

// FetchExchangeRates makes an HTTP GET request to the provided API URL and returns the exchange rates.
func FetchExchangeRates(apiURL string) ([]models.CurrencyInfo, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var currencyInfos []models.CurrencyInfo
	err = json.Unmarshal(body, &currencyInfos)
	if err != nil {
		return nil, err
	}

	return currencyInfos, nil
}
