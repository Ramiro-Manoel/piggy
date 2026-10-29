package memory

import (
	"fmt"
	"slices"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type cardTransactionRepository struct {
	transactions []transaction.CardTransaction
}

func NewCardTransactionRepository() *cardTransactionRepository {
	return &cardTransactionRepository{transactions: make([]transaction.CardTransaction, 0)}
}

func (r *cardTransactionRepository) Create(t transaction.CardTransaction) error {
	r.transactions = append(r.transactions, t)
	return nil
}

func (r *cardTransactionRepository) Read(id string) (transaction.CardTransaction, error) {
	for i := range r.transactions {
		if r.transactions[i].ID == id {
			return r.transactions[i], nil
		}
	}
	return transaction.CardTransaction{}, fmt.Errorf("transaction with id %s not found", id)
}

func (r *cardTransactionRepository) List(f transaction.CardFilter) ([]transaction.CardTransaction, error) {
	transactions := []transaction.CardTransaction{}
	for _, t := range r.transactions {
		if len(f.InvoiceIDs) > 0 && !slices.Contains(f.InvoiceIDs, t.InvoiceID) {
			continue
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
