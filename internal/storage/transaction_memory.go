package storage

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/domain"
)

type TransactionRepository interface {
	Save(t domain.Transaction) error
	Read(id string) (domain.Transaction, error)
	List() []domain.Transaction
}

type InMemoryTransactionRepository struct {
	transactions []domain.Transaction
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{transactions: make([]domain.Transaction, 0)}
}

func (r *InMemoryTransactionRepository) Save(t domain.Transaction) error {
	r.transactions = append(r.transactions, t)
	return nil
}

func (r *InMemoryTransactionRepository) Read(id string) (domain.Transaction, error) {
	for i := range r.transactions {
		if r.transactions[i].ID == id {
			return r.transactions[i], nil
		}
	}
	return domain.Transaction{}, fmt.Errorf("transaction with id %s not found", id)
}

func (r *InMemoryTransactionRepository) List() []domain.Transaction {
	transactions := append([]domain.Transaction{}, r.transactions...)
	return transactions
}
