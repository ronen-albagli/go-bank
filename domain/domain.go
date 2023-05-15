package domain

import "time"

type Transaction struct {
	ID          string
	Amount      float64
	Description string
	CreatedAt   time.Time
}

type TransactionRepository interface {
	Save(transaction *Transaction) error
}
