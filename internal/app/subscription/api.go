package subscription

import (
	"fmt"
	"log"

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

// GetAllEmails retrieves all email addresses from the database.
func GetAllEmails() ([]string, error) {
	var emails []string

	// Prepare the query to select all emails
	query := "SELECT email FROM emails"
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Printf("Error querying emails: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and append each email to the slice
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			log.Printf("Error scanning email: %v", err)
			return nil, err
		}
		emails = append(emails, email)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return emails, nil
}
