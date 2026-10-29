package handler

import (
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func (h *Handler) listAccountTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.accountTransactionSvc.List(transaction.AccountFilter{
		AccountIDs: []string{r.PathValue("accountID")},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, transactions)
}

func (h *Handler) syncAccountTransactions(w http.ResponseWriter, r *http.Request) {
	accountID := r.PathValue("accountID")

	err := h.accountTransactionSvc.Sync(accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
