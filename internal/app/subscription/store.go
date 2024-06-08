package subscription

import (
	"fmt"

	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
)

func addEmail(email string) error {
	stmt, err := config.DB.Prepare("INSERT INTO emails(email) VALUES(?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(email)
	if err != nil {
		return err
	}

	return nil
}


func fetchEmails() ([]string, error) {
	var emails []string
	query := "SELECT email FROM emails"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying emails: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			return nil, fmt.Errorf("error scanning email: %w", err)
		}
		emails = append(emails, email)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return emails, nil
}


func deleteEmail(email string) error {
	stmt, err := config.DB.Prepare("DELETE FROM emails WHERE email = ?")
	if err != nil {
		return fmt.Errorf("error preparing delete statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(email)
	if err != nil {
		return fmt.Errorf("error executing delete statement: %w", err)
	}

	// Check if the email was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, email may not exist")
	}

	return nil
}
