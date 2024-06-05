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
	mockResponse := []models.CurrencyInfo{
		{
			Txt:          "Австралійський долар",
			Rate:         1.2,
			Cc:           "AUD",
			ExchangeDate: "05.06.2024",
		},
		{
			Txt:          "Канадський долар",
			Rate:         39.2,
			Cc:           "CAD",
			ExchangeDate: "05.06.2024",
		},
	}

	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Call the function with the mock server URL
	currencyInfos, err := FetchExchangeRates(server.URL)
	if err != nil {
		t.Fatalf("FetchExchangeRates returned an error: %v", err)
	}

	// Assert that the response matches the mock response
	for i, info := range currencyInfos {
		if info.Txt != mockResponse[i].Txt {
			t.Errorf("Expected Txt to be %v, got %v", mockResponse[i].Txt, info.Txt)
		}
		if info.Rate != mockResponse[i].Rate {
			t.Errorf("Expected Rate to be %v, got %v", mockResponse[i].Rate, info.Rate)
		}
		if info.Cc != mockResponse[i].Cc {
			t.Errorf("Expected Cc to be %v, got %v", mockResponse[i].Cc, info.Cc)
		}
		if info.ExchangeDate != mockResponse[i].ExchangeDate {
			t.Errorf("Expected ExchangeDate to be %v, got %v", mockResponse[i].ExchangeDate, info.ExchangeDate)
		}
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
