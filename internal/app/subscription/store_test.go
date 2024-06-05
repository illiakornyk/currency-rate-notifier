package subscription

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/illiakornyk/currency-rate-notifier/internal/app/config"
)

func TestAddEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	config.DB = db

	mock.ExpectPrepare("INSERT INTO emails").
		ExpectExec().
		WithArgs("test@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := addEmail("test@example.com"); err != nil {
		t.Errorf("error was not expected while adding email: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
