package memory

import (
	"fmt"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

var _ transaction.Repository = (*InMemoryTransactionRepository)(nil)

type InMemoryTransactionRepository struct {
	transactions []transaction.Transaction
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{transactions: make([]transaction.Transaction, 0)}
}

func (r *InMemoryTransactionRepository) Save(t transaction.Transaction) error {
	r.transactions = append(r.transactions, t)
	return nil
}

func (r *InMemoryTransactionRepository) Read(id string) (transaction.Transaction, error) {
	for i := range r.transactions {
		if r.transactions[i].ID == id {
			return r.transactions[i], nil
		}
	}
	return transaction.Transaction{}, fmt.Errorf("transaction with id %s not found", id)
}

func (r *InMemoryTransactionRepository) List() []transaction.Transaction {
	transactions := append([]transaction.Transaction{}, r.transactions...)
	return transactions
}
