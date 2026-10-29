package invoice

import (
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/google/uuid"
)

type Status string

const (
	StatusOpen   Status = "open"
	StatusClosed Status = "closed"
	StatusPaid   Status = "paid"
)

type Invoice struct {
	ID        string
	Ref       external.Reference
	CardID    string
	CloseDate *time.Time
	PayDate   time.Time
	Amount    int64
}

func New(cardID string, closeDate *time.Time, dueDate time.Time) Invoice {
	return Invoice{
		ID:        uuid.New().String(),
		CardID:    cardID,
		CloseDate: closeDate,
		PayDate:   dueDate,
		Amount:    0,
	}
}

func (i Invoice) Status(now time.Time) Status {
	if !now.Before(i.PayDate.AddDate(0, 0, 1)) {
		return StatusPaid
	}
	if i.CloseDate != nil && !now.Before(*i.CloseDate) {
		return StatusClosed
	}
	return StatusOpen
}
