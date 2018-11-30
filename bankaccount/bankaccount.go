package bankaccount

import "database/sql"

type BankAccount struct {
	ID        int
	UserID    string
	AccountID int
	Name      string
	Balance   float64
}

type Manager struct {
	DB *sql.DB
}
