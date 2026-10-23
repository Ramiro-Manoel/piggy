package memory

import (
	"fmt"
	"slices"

	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/Ramiro-Manoel/piggy/internal/invoice"
)

type invoiceRepository struct {
	invoices []invoice.Invoice
}

func NewInvoiceRepository() *invoiceRepository {
	return &invoiceRepository{invoices: make([]invoice.Invoice, 0)}
}

func (r *invoiceRepository) Create(i invoice.Invoice) error {
	r.invoices = append(r.invoices, i)
	return nil
}

func (r *invoiceRepository) Update(i invoice.Invoice) error {
	for idx := range r.invoices {
		if r.invoices[idx].ID == i.ID {
			r.invoices[idx] = i
			return nil
		}
	}
	return invoice.ErrNotFound
}

func (r *invoiceRepository) ReadByRef(ref external.Reference) (invoice.Invoice, error) {
	for _, i := range r.invoices {
		if i.Ref == ref {
			return i, nil
		}
	}
	return invoice.Invoice{}, invoice.ErrNotFound
}

func (r *invoiceRepository) Read(id string) (invoice.Invoice, error) {
	for i := range r.invoices {
		if r.invoices[i].ID == id {
			return r.invoices[i], nil
		}
	}
	return invoice.Invoice{}, fmt.Errorf("category with id %s not found", id)
}

func (r *invoiceRepository) List(f invoice.Filter) ([]invoice.Invoice, error) {
	invoices := []invoice.Invoice{}
	for _, i := range r.invoices {
		if matches(i, f) {
			invoices = append(invoices, i)
		}
	}
	return invoices, nil
}

func matches(i invoice.Invoice, f invoice.Filter) bool {
	if len(f.CardIDs) > 0 && !slices.Contains(f.CardIDs, i.CardID) {
		return false
	}

	if f.CloseDate.Start != nil && i.CloseDate.Before(*f.CloseDate.Start) {
		return false
	}
	if f.CloseDate.End != nil && i.CloseDate.After(*f.CloseDate.End) {
		return false
	}

	return true

}
