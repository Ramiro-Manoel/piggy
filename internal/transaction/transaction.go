package transaction

import (
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/external"
)

type Transaction struct {
	ID          string
	Ref         external.Reference
	Description string
	Amount      int64
	Date        time.Time
	CategoryID  *string
	AccountID   string
}
