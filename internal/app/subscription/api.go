package subscription

import (
	"fmt"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
)

// InsertEmail inserts the provided email into the database.
func InsertEmail(email string) error {
	// Prepare the insert statement
	stmt, err := config.DB.Prepare("INSERT INTO emails(email) VALUES(?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	// Execute the statement with the provided email
	_, err = stmt.Exec(email)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}
