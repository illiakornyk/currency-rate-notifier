package subscription

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/models"
)

// AddSubscriber inserts the provided email into the database.
func AddSubscriber(email string) error {
	err := addEmail(email)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf(models.ErrEmailAlreadySubscribed)
		}
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}

// RetrieveSubscribers retrieves all email addresses from the database.
func RetrieveSubscribers() ([]string, error) {
	emails, err := fetchEmails()
	if err != nil {
		log.Printf("Error retrieving subscribers: %v", err)
		return nil, err
	}
	return emails, nil
}
