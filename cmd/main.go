package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/illiakornyk/currency-rate-notifier/cmd/api"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
	"github.com/illiakornyk/currency-rate-notifier/internal/scheduler"
)



func main() {
	config.Init()

	api.RunApiServer()

	scheduler.SetupCronJobs()


	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
