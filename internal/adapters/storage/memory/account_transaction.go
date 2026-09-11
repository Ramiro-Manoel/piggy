package memory

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type AccountTransactionRepository struct {
	transactions []transaction.AccountTransaction
}

func NewAccountTransactionRepository() *AccountTransactionRepository {
	return &AccountTransactionRepository{transactions: make([]transaction.AccountTransaction, 0)}
}

func (r *AccountTransactionRepository) Save(t transaction.AccountTransaction) error {
	r.transactions = append(r.transactions, t)
	return nil
}

func (r *AccountTransactionRepository) Read(id string) (transaction.AccountTransaction, error) {
	for i := range r.transactions {
		if r.transactions[i].ID == id {
			return r.transactions[i], nil
		}
	}
	return transaction.AccountTransaction{}, fmt.Errorf("transaction with id %s not found", id)
}

func (r *AccountTransactionRepository) List() []transaction.AccountTransaction {
	transactions := append([]transaction.AccountTransaction{}, r.transactions...)
	return transactions
}
