package scheduler

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/handlers"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/subscription"
	"github.com/illiakornyk/currency-rate-notifier/internal/email"

	"github.com/robfig/cron/v3"
)

func SetupCronJobs() {
	// Create a new cron instance with Kyiv's timezone
	loc, _ := time.LoadLocation("Europe/Kyiv")
	c := cron.New(cron.WithLocation(loc), cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))

	// Schedule the job to run every day at 8 AM Kyiv time
_, err := c.AddFunc("0 8 * * *", func() {
		apiURL, err := handlers.ConstructAPIURL()
		if err != nil {
			log.Fatal(err)
			return
		}

		exchangeRates, err := handlers.FetchExchangeRateData(apiURL)
		if err != nil {
			log.Println("Error fetching exchange rate:", err)
			return
		}

		exchangeDate := exchangeRates[0].ExchangeDate

		// Filter for EUR and USD rates
		var eurRate, usdRate float64
		for _, rate := range exchangeRates {
			if rate.Cc == "EUR" {
				eurRate = rate.Rate
			} else if rate.Cc == "USD" {
				usdRate = rate.Rate
			}
		}

		// Retrieve all subscribed email addresses
		emails, err := subscription.RetrieveSubscribers()
		if err != nil {
			log.Println("Error retrieving emails:", err)
			return
		}

		// Send an email to each address with EUR and USD rates
		for _, receiverEmail := range emails {
			emailSender := os.Getenv("GMAIL_SMTP_EMAIL")
			emailSenderPassword := os.Getenv("GMAIL_SMTP_PASSWORD")
			subject := "Daily Exchange Rate"

			body := fmt.Sprintf("Today's exchange rates for %s are:\nEUR to UAH: %.2f\nUSD to UAH: %.2f", exchangeDate, eurRate, usdRate)

			err := email.SendEmail(emailSender, emailSenderPassword, receiverEmail, subject, body)
			if err != nil {
				log.Println("Error sending email to", receiverEmail, ":", err)
			}
		}
	})

	if err != nil {
		log.Fatal("Error adding func to cron:", err)
	}

	c.Start()
}
