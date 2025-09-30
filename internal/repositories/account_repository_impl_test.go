package repositories

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/gregoryAlvim/gobank/internal/models"
)

func TestPsqlAccountRepository_CreateNaturalPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &PsqlAccountRepository{DB: db}

	person := &models.NaturalPerson{
		MonthlyIncome: 5000,
		Age:           30,
		FullName:      "John Doe",
		PhoneNumber:   "123456789",
		Email:         "john.doe@example.com",
		Category:      "standard",
		Balance:       1000,
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(`INSERT INTO natural_person`).
		WithArgs(person.MonthlyIncome, person.Age, person.FullName, person.PhoneNumber, person.Email, person.Category, person.Balance).
		WillReturnRows(rows)

	err = repo.CreateNaturalPerson(person)

	assert.NoError(t, err)
	assert.Equal(t, 1, person.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPsqlAccountRepository_CreateLegalPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &PsqlAccountRepository{DB: db}

	person := &models.LegalPerson{
		AnnualRevenue:  100000,
		Age:            5,
		TradeName:      "ABC Inc.",
		PhoneNumber:    "987654321",
		CorporateEmail: "contact@abcinc.com",
		Category:       "business",
		Balance:        50000,
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(`INSERT INTO legal_person`).
		WithArgs(person.AnnualRevenue, person.Age, person.TradeName, person.PhoneNumber, person.CorporateEmail, person.Category, person.Balance).
		WillReturnRows(rows)

	err = repo.CreateLegalPerson(person)

	assert.NoError(t, err)
	assert.Equal(t, 1, person.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPsqlAccountRepository_GetAccountBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &PsqlAccountRepository{DB: db}

	// Test for natural person
	rows := sqlmock.NewRows([]string{"balance"}).AddRow(123.45)
	mock.ExpectQuery("SELECT balance FROM natural_person").WithArgs(1).WillReturnRows(rows)
	balance, err := repo.GetAccountBalance(1, "natural")
	assert.NoError(t, err)
	assert.Equal(t, 123.45, balance)

	// Test for legal person
	rows = sqlmock.NewRows([]string{"balance"}).AddRow(543.21)
	mock.ExpectQuery("SELECT balance FROM legal_person").WithArgs(1).WillReturnRows(rows)
	balance, err = repo.GetAccountBalance(1, "legal")
	assert.NoError(t, err)
	assert.Equal(t, 543.21, balance)

	// Test for not found
	mock.ExpectQuery("SELECT balance FROM natural_person").WithArgs(1).WillReturnError(sql.ErrNoRows)
	_, err = repo.GetAccountBalance(1, "natural")
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPsqlAccountRepository_UpdateAccountBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &PsqlAccountRepository{DB: db}

	mock.ExpectExec("UPDATE natural_person").WithArgs(200.0, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.UpdateAccountBalance(1, 200.0, "natural")
	assert.NoError(t, err)

	mock.ExpectExec("UPDATE legal_person").WithArgs(1000.0, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.UpdateAccountBalance(1, 1000.0, "legal")
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPsqlAccountRepository_DeleteAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &PsqlAccountRepository{DB: db}

	mock.ExpectExec("DELETE FROM natural_person").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.DeleteAccount(1, "natural")
	assert.NoError(t, err)

	mock.ExpectExec("DELETE FROM legal_person").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = repo.DeleteAccount(1, "legal")
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPsqlAccountRepository_TransferTx(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &PsqlAccountRepository{DB: db}

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT balance FROM natural_person").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(500.0))
	mock.ExpectQuery("SELECT balance FROM legal_person").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(1000.0))
	mock.ExpectExec("UPDATE natural_person").WithArgs(400.0, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE legal_person").WithArgs(1100.0, 2).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.TransferTx(1, 2, 100.0, "natural", "legal")
	assert.NoError(t, err)

	// Test insufficient funds
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT balance FROM natural_person").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(50.0))
	mock.ExpectRollback()

	err = repo.TransferTx(1, 2, 100.0, "natural", "legal")
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
