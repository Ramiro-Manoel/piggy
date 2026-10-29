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

func (s service) FindOrCreate(c card.Card, month time.Time) (Invoice, error) {
	invoices, err := s.repo.List(Filter{
		CardIDs: []string{c.ID},
		PayDate: monthRange(month),
	})
	if err != nil {
		return Invoice{}, fmt.Errorf("FindOrCreate: %w", err)
	}
	if len(invoices) > 0 {
		return invoices[0], nil
	}

	if c.DueDay == nil {
		return Invoice{}, fmt.Errorf("FindOrCreate: card %s has no due day", c.ID)
	}
	dueDate := time.Date(month.Year(), month.Month(), *c.DueDay, 0, 0, 0, 0, time.UTC)

	var closeDate *time.Time
	if c.ClosingDay != nil {
		d := time.Date(month.Year(), month.Month(), *c.ClosingDay, 0, 0, 0, 0, time.UTC)
		if *c.ClosingDay > *c.DueDay {
			d = d.AddDate(0, -1, 0)
		}
		closeDate = &d
	}

	invoice := New(c.ID, closeDate, dueDate)
	if err := s.repo.Create(invoice); err != nil {
		return Invoice{}, fmt.Errorf("FindOrCreate: %w", err)
	}
	return invoice, nil
}
