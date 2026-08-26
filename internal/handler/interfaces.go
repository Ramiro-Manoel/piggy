package handler

import (
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type transactionService interface {
	Create(transaction.Transaction) error
	Read(string) (transaction.Transaction, error)
	List() []transaction.Transaction
}

type categoryService interface {
	Create(category.Category) error
	Read(string) (category.Category, error)
	List() []category.Category
}
