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
	PaidAt    *time.Time // nil enquanto a fatura não foi paga
}

func New(cardID string, closeDate time.Time, dueDate time.Time) Invoice {
	return Invoice{
		ID:        uuid.New().String(),
		CardID:    cardID,
		CloseDate: closeDate,
		PayDate:   dueDate,
		Amount:    0,
	}
}

// Status calcula a situação da fatura em relação a now. Não é guardado no
// banco porque open/closed depende só da data e ficaria desatualizado.
func (i Invoice) Status(now time.Time) Status {
	if i.PaidAt != nil {
		return StatusPaid
	}
	if now.Before(i.CloseDate) {
		return StatusOpen
	}
	return StatusClosed
}
