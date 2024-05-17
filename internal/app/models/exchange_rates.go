package models

// ExchangeRatesResponse represents the response structure from the exchange rates API
type ExchangeRatesResponse struct {
    Success   bool              `json:"success"`
    Timestamp int64             `json:"timestamp"`
    Base      string            `json:"base"`
    Date      string            `json:"date"`
    Rates     map[string]float64 `json:"rates"`
}
