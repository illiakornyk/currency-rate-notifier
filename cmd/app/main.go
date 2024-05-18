package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/handlers"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/subscription"
)



func main() {

	config.Init()

	http.HandleFunc("/rate", handlers.RatesHandler)

	err := subscription.InsertEmail("myMail@example.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Email inserted successfully!")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
