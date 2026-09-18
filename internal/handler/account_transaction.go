package handler

import (
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func (h *Handler) createAccountTransaction(w http.ResponseWriter, r *http.Request) {
	t, err := decode[transaction.AccountTransaction](w, r)
	if err != nil {
		return
	}

	err = h.accountTransactionSvc.Create(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) listAccountTransactions(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.accountTransactionSvc.List())
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
