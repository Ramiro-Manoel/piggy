package card

import "github.com/Ramiro-Manoel/piggy/internal/external"

type Card struct {
	ID             string
	Ref            external.Reference
	Name           string
	Brand          string
	CreditLimit    int64
	AvailableLimit int64
	ClosingDay     int
	DueDay         int
	AccountID      *string
}
