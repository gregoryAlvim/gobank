package services

import (
	"encoding/json"
	"errors"

	"github.com/gregoryAlvim/gobank/models"
	"github.com/gregoryAlvim/gobank/repositories"
)

type AccountService struct {
	repo repositories.AccountRepository
}

func NewAccountService(repo repositories.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(accountType string, data []byte) error {
	switch accountType {
	case "natural":
		var person models.NaturalPerson
		if err := json.Unmarshal(data, &person); err != nil {
			return err
		}
		return s.repo.CreateNaturalPerson(&person)
	case "legal":
		var person models.LegalPerson
		if err := json.Unmarshal(data, &person); err != nil {
			return err
		}
		return s.repo.CreateLegalPerson(&person)
	default:
		return errors.New("invalid account type")
	}
}

func (s *AccountService) GetBalance(accountID int, accountType string) (float64, error) {
	return s.repo.GetAccountBalance(accountID, accountType)
}

func (s *AccountService) Deposit(accountID int, amount float64, accountType string) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}

	currentBalance, err := s.repo.GetAccountBalance(accountID, accountType)
	if err != nil {
		return err
	}

	newBalance := currentBalance + amount
	return s.repo.UpdateAccountBalance(accountID, newBalance, accountType)
}

func (s *AccountService) Withdraw(accountID int, amount float64, accountType string) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	currentBalance, err := s.repo.GetAccountBalance(accountID, accountType)
	if err != nil {
		return err
	}

	if currentBalance < amount {
		return errors.New("insufficient funds")
	}

	newBalance := currentBalance - amount
	return s.repo.UpdateAccountBalance(accountID, newBalance, accountType)
}

// Transfer performs the money transfer between two accounts within a transaction.
func (s *AccountService) Transfer(fromID, toID int, amount float64, fromType, toType string) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	// The actual withdrawal and deposit will be handled by the repository
	// within a single database transaction to ensure atomicity.
	return s.repo.TransferTx(fromID, toID, amount, fromType, toType)
}

func (s *AccountService) CloseAccount(accountID int, accountType string) error {
	return s.repo.DeleteAccount(accountID, accountType)
}
