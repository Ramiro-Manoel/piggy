package handler

import "net/http"

func (h *Handler) listCardInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := h.invoiceSvc.List(r.PathValue("cardID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, invoices)
}

func (h *Handler) listInvoiceTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.cardTransactionSvc.ListByInvoice(r.PathValue("invoiceID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, transactions)
}
