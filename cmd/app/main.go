package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/handlers"
	"github.com/illiakornyk/currency-rate-notifier/internal/scheduler"
)



func main() {
	config.Init()

	http.HandleFunc("/rate", handlers.RatesHandler)
	http.HandleFunc("/subscribe", handlers.SubscribeHandler)

	scheduler.SetupCronJobs()


	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
