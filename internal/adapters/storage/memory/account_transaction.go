package memory

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type accountTransactionRepository struct {
	transactions []transaction.AccountTransaction
}

func NewAccountTransactionRepository() *accountTransactionRepository {
	return &accountTransactionRepository{transactions: make([]transaction.AccountTransaction, 0)}
}

func (r *accountTransactionRepository) Save(t transaction.AccountTransaction) error {
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

func (r *accountTransactionRepository) List() []transaction.AccountTransaction {
	transactions := append([]transaction.AccountTransaction{}, r.transactions...)
	return transactions
}
