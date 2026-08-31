package transaction

import (
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/provider"
)

type Transaction struct {
	ID          string
	Ref         provider.Ref
	Description string
	Amount      int64
	Date        time.Time
	CategoryID  *string
	AccountID   string
}
