package invoice

import (
	"time"

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
	CardID    string
	CloseDate time.Time
	PayDate   time.Time
	Amount    int64
	Status    Status
}

func New(cardID string, closeDate time.Time, dueDate time.Time) Invoice {
	status := StatusOpen
	if closeDate.Before(time.Now()) {
		status = StatusClosed
	}

	return Invoice{
		ID:        uuid.New().String(),
		CardID:    cardID,
		CloseDate: closeDate,
		PayDate:   dueDate,
		Amount:    0,
		Status:    status,
	}
}
