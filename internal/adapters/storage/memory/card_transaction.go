package memory

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type CardTransactionRepository struct {
	transactions []transaction.CardTransaction
}

func NewCardTransactionRepository() *CardTransactionRepository {
	return &CardTransactionRepository{transactions: make([]transaction.CardTransaction, 0)}
}

func (r *CardTransactionRepository) Save(t transaction.CardTransaction) error {
	r.transactions = append(r.transactions, t)
	return nil
}

func (r *CardTransactionRepository) Read(id string) (transaction.CardTransaction, error) {
	for i := range r.transactions {
		if r.transactions[i].ID == id {
			return r.transactions[i], nil
		}
	}
	return transaction.CardTransaction{}, fmt.Errorf("transaction with id %s not found", id)
}

func (r *CardTransactionRepository) List() []transaction.CardTransaction {
	transactions := append([]transaction.CardTransaction{}, r.transactions...)
	return transactions
}
