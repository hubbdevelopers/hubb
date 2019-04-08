package repositories

import (
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestExampleSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO \"users\"")

	gdb, err := gorm.Open("sqlmock", db)
	if err != nil {
		log.Fatal(err)
	}

	repo := NewUserRepository(gdb)
	repo.Create("testuser")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
