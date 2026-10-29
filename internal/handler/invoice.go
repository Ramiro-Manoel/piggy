package handler

import (
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/invoice"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func (h *Handler) listCardInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := h.invoiceSvc.List(invoice.Filter{
		CardIDs: []string{r.PathValue("cardID")},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, invoices)
}

func (h *Handler) listInvoiceTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.cardTransactionSvc.List(transaction.CardFilter{
		InvoiceIDs: []string{r.PathValue("invoiceID")},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, transactions)
}
