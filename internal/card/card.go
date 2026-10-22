package card

import "github.com/Ramiro-Manoel/piggy/internal/external"

type Card struct {
	ID             string
	Ref            external.Reference
	Name           string
	Brand          string
	CreditLimit    int64
	AvailableLimit int64
	ClosingDay     *int // nil quando o banco não informa
	DueDay         *int // nil quando o banco não informa
	AccountID      *string
}
