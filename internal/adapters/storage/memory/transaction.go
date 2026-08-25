package memory

import (
	"fmt"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

var _ transaction.Repository = (*transactionRepository)(nil)

type transactionRepository struct {
	transactions []transaction.Transaction
}

func NewtransactionRepository() *transactionRepository {
	return &transactionRepository{transactions: make([]transaction.Transaction, 0)}
}

func (r *transactionRepository) Save(t transaction.Transaction) error {
	r.transactions = append(r.transactions, t)
	return nil
}

func (r *transactionRepository) Read(id string) (transaction.Transaction, error) {
	for i := range r.transactions {
		if r.transactions[i].ID == id {
			return r.transactions[i], nil
		}
	}
	return transaction.Transaction{}, fmt.Errorf("transaction with id %s not found", id)
}

func (r *transactionRepository) List() []transaction.Transaction {
	transactions := append([]transaction.Transaction{}, r.transactions...)
	return transactions
}
