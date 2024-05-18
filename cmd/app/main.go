package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/handlers"
)



func main() {

	config.Init()

	http.HandleFunc("/rate", handlers.RatesHandler)
	http.HandleFunc("/subscribe", handlers.SubscribeHandler)


	fmt.Println("Email inserted successfully!")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
