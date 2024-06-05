package api

import (
	"net/http"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/handlers"
)



func RunApiServer()  {
	http.HandleFunc("/rate", handlers.RatesHandler)
	http.HandleFunc("/subscribe", handlers.SubscribeHandler)
}
