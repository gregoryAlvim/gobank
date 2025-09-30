package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mymodule/database"
	"mymodule/models"
)

type PsqlAccountRepository struct {
	DB *sql.DB
}

func NewPsqlAccountRepository() *PsqlAccountRepository {
	return &PsqlAccountRepository{DB: database.DB}
}

func (r *PsqlAccountRepository) CreateNaturalPerson(person *models.NaturalPerson) error {
	query := `INSERT INTO natural_person (monthly_income, age, full_name, phone_number, email, category, balance)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := r.DB.QueryRow(query, person.MonthlyIncome, person.Age, person.FullName, person.PhoneNumber, person.Email, person.Category, person.Balance).Scan(&person.ID)
	return err
}

func (r *PsqlAccountRepository) CreateLegalPerson(person *models.LegalPerson) error {
	query := `INSERT INTO legal_person (annual_revenue, age, trade_name, phone_number, corporate_email, category, balance)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := r.DB.QueryRow(query, person.AnnualRevenue, person.Age, person.TradeName, person.PhoneNumber, person.CorporateEmail, person.Category, person.Balance).Scan(&person.ID)
	return err
}

func (r *PsqlAccountRepository) GetAccountBalance(accountID int, accountType string) (float64, error) {
	var balance float64
	var query string

	switch accountType {
	case "natural":
		query = "SELECT balance FROM natural_person WHERE id = $1"
	case "legal":
		query = "SELECT balance FROM legal_person WHERE id = $1"
	default:
		return 0, errors.New("invalid account type")
	}

	err := r.DB.QueryRow(query, accountID).Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("account not found")
		}
		return 0, err
	}
	return balance, nil
}

func (r *PsqlAccountRepository) UpdateAccountBalance(accountID int, newBalance float64, accountType string) error {
	var query string
	switch accountType {
	case "natural":
		query = "UPDATE natural_person SET balance = $1 WHERE id = $2"
	case "legal":
		query = "UPDATE legal_person SET balance = $1 WHERE id = $2"
	default:
		return errors.New("invalid account type")
	}

	_, err := r.DB.Exec(query, newBalance, accountID)
	return err
}

func (r *PsqlAccountRepository) DeleteAccount(accountID int, accountType string) error {
	var query string
	switch accountType {
	case "natural":
		query = "DELETE FROM natural_person WHERE id = $1"
	case "legal":
		query = "DELETE FROM legal_person WHERE id = $1"
	default:
		return errors.New("invalid account type")
	}

	_, err := r.DB.Exec(query, accountID)
	return err
}

func (r *PsqlAccountRepository) TransferTx(fromID, toID int, amount float64, fromType, toType string) error {
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() // Rollback on any error.

	// 1. Get and check fromAccount's balance
	fromBalance, err := r.getAccountBalanceTx(tx, fromID, fromType)
	if err != nil {
		return err
	}
	if fromBalance < amount {
		return errors.New("insufficient funds")
	}

	// 2. Get toAccount's balance
	toBalance, err := r.getAccountBalanceTx(tx, toID, toType)
	if err != nil {
		return err
	}

	// 3. Update balances
	if err := r.updateAccountBalanceTx(tx, fromID, fromBalance-amount, fromType); err != nil {
		return err
	}
	if err := r.updateAccountBalanceTx(tx, toID, toBalance+amount, toType); err != nil {
		return err
	}

	// 4. Commit the transaction
	return tx.Commit()
}

// Helper functions to be used within a transaction
func (r *PsqlAccountRepository) getAccountBalanceTx(tx *sql.Tx, accountID int, accountType string) (float64, error) {
	var balance float64
	var query string

	switch accountType {
	case "natural":
		query = "SELECT balance FROM natural_person WHERE id = $1 FOR UPDATE"
	case "legal":
		query = "SELECT balance FROM legal_person WHERE id = $1 FOR UPDATE"
	default:
		return 0, errors.New("invalid account type")
	}

	err := tx.QueryRow(query, accountID).Scan(&balance)
	return balance, err
}

func (r *PsqlAccountRepository) updateAccountBalanceTx(tx *sql.Tx, accountID int, newBalance float64, accountType string) error {
	var query string
	switch accountType {
	case "natural":
		query = "UPDATE natural_person SET balance = $1 WHERE id = $2"
	case "legal":
		query = "UPDATE legal_person SET balance = $1 WHERE id = $2"
	default:
		return errors.New("invalid account type")
	}

	_, err := tx.Exec(query, newBalance, accountID)
	return err
}
