package transaction

import (
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/card"
	"github.com/Ramiro-Manoel/piggy/internal/invoice"
)

type cardReader interface {
	Read(id string) (card.Card, error)
}

type invoiceService interface {
	FindOrCreate(c card.Card, transactionDate time.Time) (invoice.Invoice, error)
}
