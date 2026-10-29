package handler

import (
	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/card"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/invoice"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type accountTransactionService interface {
	ListByAccount(accountID string) ([]transaction.AccountTransaction, error)
	Sync(string) error
}

type cardTransactionService interface {
	ListByInvoice(invoiceID string) ([]transaction.CardTransaction, error)
	Sync(string) error
}

type categoryService interface {
	Create(category.Category) error
	List() []category.Category
}

type accountService interface {
	List() []account.Account
	Sync(string) error
}

type cardService interface {
	List() []card.Card
	Sync(string) error
}

type invoiceService interface {
	List(cardID string) ([]invoice.Invoice, error)
}
