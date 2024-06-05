package models

// ExchangeRatesResponse represents the response structure from the exchange rates API
type CurrencyInfo struct {
	// R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	Cc           string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}
