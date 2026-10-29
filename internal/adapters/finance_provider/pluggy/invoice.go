package pluggy

import (
	"math"
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/Ramiro-Manoel/piggy/internal/invoice"
)

func (c *client) FetchInvoices(accountID string) ([]invoice.Invoice, error) {
	billsResp, err := c.fetchBills(accountID)
	if err != nil {
		return []invoice.Invoice{}, err
	}

	return toInvoices(billsResp.Results)
}

func toInvoice(bill pluggyBill) (invoice.Invoice, error) {
	closeDate, err := time.Parse(time.RFC3339Nano, bill.BillClosingDate)
	if err != nil {
		return invoice.Invoice{}, err
	}
	dueDate, err := time.Parse(time.RFC3339Nano, bill.DueDate)
	if err != nil {
		return invoice.Invoice{}, err
	}

	return invoice.Invoice{
		Ref: external.Reference{
			ExternalID: bill.ID,
			Source:     source},
		CloseDate: &closeDate,
		PayDate:   dueDate,
		Amount:    int64(math.Round(bill.TotalAmount * 100)),
	}, nil
}

func toInvoices(bills []pluggyBill) ([]invoice.Invoice, error) {
	var invoices []invoice.Invoice
	for _, bill := range bills {
		i, err := toInvoice(bill)
		if err != nil {
			return []invoice.Invoice{}, err
		}
		invoices = append(invoices, i)
	}
	return invoices, nil
}
