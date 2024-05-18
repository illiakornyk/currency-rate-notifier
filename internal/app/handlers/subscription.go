package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/subscription"
)

// EmailRequest represents the JSON structure for the email request
type EmailRequest struct {
	Email string `json:"email"`
}

// SubscribeHandler handles requests for the /subscribe route
func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON body
	var req EmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the email into the database
	err = subscription.InsertEmail(req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Subscription successful"))
}
