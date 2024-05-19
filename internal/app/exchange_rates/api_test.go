package exchange_rates

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/models"
)

func TestFetchExchangeRates(t *testing.T) {
	// Define a mock response
	mockResponse := &models.ExchangeRatesResponse{
		Success: true,
		Rates: map[string]float64{
			"USD": 1.2,
			"UAH": 39.2,
		},
	}

	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Call the function with the mock server URL
	exchangeRates, err := FetchExchangeRates(server.URL)
	if err != nil {
		t.Fatalf("FetchExchangeRates returned an error: %v", err)
	}

	// Assert that the response matches the mock response
	if exchangeRates.Success != mockResponse.Success {
		t.Errorf("Expected Success to be %v, got %v", mockResponse.Success, exchangeRates.Success)
	}
	if exchangeRates.Rates["USD"] != mockResponse.Rates["USD"] {
		t.Errorf("Expected USD rate to be %v, got %v", mockResponse.Rates["USD"], exchangeRates.Rates["USD"])
	}
	if exchangeRates.Rates["UAH"] != mockResponse.Rates["UAH"] {
		t.Errorf("Expected UAH rate to be %v, got %v", mockResponse.Rates["UAH"], exchangeRates.Rates["UAH"])
	}
}


func TestFetchExchangeRates_APIFailure(t *testing.T) {
	// Create a mock server that returns an HTTP error status
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Call the function with the mock server URL
	_, err := FetchExchangeRates(server.URL)
	if err == nil {
		t.Error("Expected an error for API failure, got nil")
	}
}


func TestFetchExchangeRates_InvalidJSON(t *testing.T) {
	// Create a mock server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid JSON"))
	}))
	defer server.Close()

	// Call the function with the mock server URL
	_, err := FetchExchangeRates(server.URL)
	if err == nil {
		t.Error("Expected an error for invalid JSON, got nil")
	}
}
