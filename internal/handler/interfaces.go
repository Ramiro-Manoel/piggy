package handler

import (
	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type accountTransactionService interface {
	Create(transaction.AccountTransaction) error
	Read(string) (transaction.AccountTransaction, error)
	List() []transaction.AccountTransaction
	Sync(string) error
}

type cardTransactionService interface {
	Create(transaction.CardTransaction) error
	Read(string) (transaction.CardTransaction, error)
	List() []transaction.CardTransaction
	Sync(string) error
}

type categoryService interface {
	Create(category.Category) error
	Read(string) (category.Category, error)
	List() []category.Category
}

type accountService interface {
	Create(account.Account) error
	Read(string) (account.Account, error)
	List() []account.Account
	Sync(string) error
}
