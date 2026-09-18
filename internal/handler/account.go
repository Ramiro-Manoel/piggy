package handler

import (
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/account"
)

func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	a, err := decode[account.Account](w, r)
	if err != nil {
		return
	}

	err = h.accountSvc.Create(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) listAccounts(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.accountSvc.List())
}

func (h *Handler) syncAccounts(w http.ResponseWriter, r *http.Request) {
	err := h.accountSvc.Sync(h.institutionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
