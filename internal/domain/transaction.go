package domain

import "time"

type Transaction struct {
	ID          string
	Description string
	Amount      int64
	Date        time.Time
}
