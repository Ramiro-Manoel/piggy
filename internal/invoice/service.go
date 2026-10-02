package invoice

import (
	"fmt"
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/card"
)

type service struct {
	repo repository
}

func NewService(r repository) *service {
	return &service{
		repo: r,
	}
}

func (s service) FindOrCreate(c card.Card, transactionDate time.Time) (Invoice, error) {
	closeDate := nextOccurrence(transactionDate, c.ClosingDay)

	invoices, err := s.repo.List(Filter{
		CloseDate: DateRange{Start: &closeDate, End: &closeDate},
		CardIDs:   []string{c.ID},
	})
	if err != nil {
		return Invoice{}, err
	}
	if len(invoices) > 0 {
		return invoices[0], nil
	}

	invoice := New(c.ID, closeDate, nextOccurrence(closeDate, c.DueDay))
	if err := s.repo.Save(invoice); err != nil {
		return Invoice{}, fmt.Errorf("FindOrCreate: %w", err)
	}
	return invoice, nil
}
