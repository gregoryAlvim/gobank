package repositories

import "github.com/gregoryAlvim/gobank/internal/models"

type AccountRepository interface {
	CreateNaturalPerson(person *models.NaturalPerson) error
	CreateLegalPerson(person *models.LegalPerson) error
	GetAccountBalance(accountID int, accountType string) (float64, error)
	UpdateAccountBalance(accountID int, newBalance float64, accountType string) error
	DeleteAccount(accountID int, accountType string) error
	TransferTx(fromID, toID int, amount float64, fromType, toType string) error
}
