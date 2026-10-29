package memory

import (
	"fmt"
	"slices"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type accountTransactionRepository struct {
	transactions []transaction.AccountTransaction
}

func NewAccountTransactionRepository() *accountTransactionRepository {
	return &accountTransactionRepository{transactions: make([]transaction.AccountTransaction, 0)}
}

func (r *accountTransactionRepository) Create(t transaction.AccountTransaction) error {
	r.transactions = append(r.transactions, t)
	return nil
}

func (r *accountTransactionRepository) Read(id string) (transaction.AccountTransaction, error) {
	for i := range r.transactions {
		if r.transactions[i].ID == id {
			return r.transactions[i], nil
		}
	}
	return transaction.AccountTransaction{}, fmt.Errorf("transaction with id %s not found", id)
}

func (r *accountTransactionRepository) List(f transaction.AccountFilter) ([]transaction.AccountTransaction, error) {
	transactions := []transaction.AccountTransaction{}
	for _, t := range r.transactions {
		if len(f.AccountIDs) > 0 && !slices.Contains(f.AccountIDs, t.AccountID) {
			continue
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
