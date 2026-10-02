package memory

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/invoice"
)

type invoiceRepository struct {
	invoices []invoice.Invoice
}

func NewInvoiceRepository() *invoiceRepository {
	return &invoiceRepository{invoices: make([]invoice.Invoice, 0)}
}

func (r *invoiceRepository) Save(i invoice.Invoice) error {
	r.invoices = append(r.invoices, i)
	return nil
}

func (r *invoiceRepository) Read(id string) (invoice.Invoice, error) {
	for i := range r.invoices {
		if r.invoices[i].ID == id {
			return r.invoices[i], nil
		}
	}
	return invoice.Invoice{}, fmt.Errorf("category with id %s not found", id)
}

func (r *invoiceRepository) List() []invoice.Invoice {
	categories := append([]invoice.Invoice{}, r.invoices...)
	return categories
}
