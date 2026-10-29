package handler

import "net/http"

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
