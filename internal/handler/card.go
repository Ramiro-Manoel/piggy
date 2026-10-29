package handler

import "net/http"

func (h *Handler) listCards(w http.ResponseWriter, r *http.Request) {
	cards, err := h.cardSvc.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, cards)
}

func (h *Handler) syncCards(w http.ResponseWriter, r *http.Request) {
	if err := h.cardSvc.Sync(h.institutionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) syncCardTransactions(w http.ResponseWriter, r *http.Request) {
	cardID := r.PathValue("cardID")
	if err := h.cardTransactionSvc.Sync(cardID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
