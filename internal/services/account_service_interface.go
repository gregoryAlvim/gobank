package services

type AccountServiceInterface interface {
	CreateAccount(accountType string, data []byte) error
	GetBalance(accountID int, accountType string) (float64, error)
	Deposit(accountID int, amount float64, accountType string) error
	Withdraw(accountID int, amount float64, accountType string) error
	Transfer(fromID, toID int, amount float64, fromType, toType string) error
	CloseAccount(accountID int, accountType string) error
}
