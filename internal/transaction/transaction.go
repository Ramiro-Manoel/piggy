package transaction

import (
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/external"
)

type transaction struct {
	ID          string
	Ref         external.Reference
	Description string
	Amount      int64
	Date        time.Time
	CategoryID  *string
}

type AccountTransaction struct {
	transaction
	AccountID string
}

type CardTransaction struct {
	transaction
	InvoiceID         string
	InvoiceMonth      *time.Time
	InstallmentNumber int
	TotalInstallments int
}
