package handler

import (
	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/card"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/invoice"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type accountTransactionService interface {
	List(transaction.AccountFilter) ([]transaction.AccountTransaction, error)
	Sync(string) error
}

type cardTransactionService interface {
	List(transaction.CardFilter) ([]transaction.CardTransaction, error)
	Sync(string) error
}

type categoryService interface {
	Create(category.Category) error
	List() ([]category.Category, error)
}

type accountService interface {
	List() ([]account.Account, error)
	Sync(string) error
}

type cardService interface {
	List() ([]card.Card, error)
	Sync(string) error
}

type invoiceService interface {
	List(invoice.Filter) ([]invoice.Invoice, error)
}
