package subscription

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
)

func TestInsertEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.DB = db

	// Expectations for database interactions
	mock.ExpectPrepare("INSERT INTO emails").
		ExpectExec().
		WithArgs("test@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function to test
	err = InsertEmail("test@example.com")
	if err != nil {
		t.Errorf("error was not expected while inserting email: %s", err)
	}

	// Make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}


func TestGetAllEmails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.DB = db

	rows := sqlmock.NewRows([]string{"email"}).
		AddRow("test@example.com").
		AddRow("another@example.com")

	mock.ExpectQuery("SELECT email FROM emails").
		WillReturnRows(rows)

	// Call the function to test
	emails, err := GetAllEmails()
	if err != nil {
		t.Errorf("error was not expected while getting emails: %s", err)
	}

	// Check the results
	expectedEmails := []string{"test@example.com", "another@example.com"}
	if !reflect.DeepEqual(emails, expectedEmails) {
		t.Errorf("expected emails to be %v, got %v", expectedEmails, emails)
	}

	// Make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
