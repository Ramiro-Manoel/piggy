package handler

import "net/http"

func (h *Handler) listAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accountSvc.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, accounts)
}

func (h *Handler) syncAccounts(w http.ResponseWriter, r *http.Request) {
	err := h.accountSvc.Sync(h.institutionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
