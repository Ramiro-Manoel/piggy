package handler

import (
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/card"
)

func (h *Handler) createCard(w http.ResponseWriter, r *http.Request) {
	c, err := decode[card.Card](w, r)
	if err != nil {
		return
	}

	if err := h.cardSvc.Create(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) listCards(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.cardSvc.List())
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
