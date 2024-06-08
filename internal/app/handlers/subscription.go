package handlers

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/models"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/subscription"
)

// EmailRequest represents the JSON structure for the email request
type EmailRequest struct {
	Email string `json:"email"`
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		subscribe(w, r)
	case http.MethodDelete:
		unsubscribe(w, r)
	default:
		http.Error(w, models.ErrMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}

// subscribe handles subscription requests
func subscribe(w http.ResponseWriter, r *http.Request) {
	// Initialize a new json.Decoder instance
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields

	// Parse the JSON body
	var req EmailRequest
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the email format
	if _, err := mail.ParseAddress(req.Email); err != nil {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Insert the email into the database
	err = subscription.AddSubscriber(req.Email)
	if err != nil {
	if err.Error() == models.ErrEmailAlreadySubscribed {
		// If the error is due to a duplicate email, send a 409 Conflict response
		http.Error(w, err.Error(), http.StatusConflict)
	} else {
		// For all other errors, send a 500 Internal Server Error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Subscription successful"))
}

// unsubscribe handles unsubscription requests
func unsubscribe(w http.ResponseWriter, r *http.Request) {
	// Initialize a new json.Decoder instance
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields

	// Parse the JSON body
	var req EmailRequest
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the email format
	if _, err := mail.ParseAddress(req.Email); err != nil {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Remove the email from the database
	err = subscription.RemoveSubscriber(req.Email)
	if err != nil {
		// Handle errors, e.g., email not found or server error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Unsubscription successful"))
}
