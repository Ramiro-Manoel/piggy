package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type Handler struct {
	transactionSvc transactionService
	categorySvc    categoryService
}

func NewHandler(transactionSvc transactionService, categorySvc categoryService) *Handler {
	return &Handler{
		transactionSvc: transactionSvc,
		categorySvc:    categorySvc,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /transactions", h.listTransactions)
	mux.HandleFunc("POST /transactions", h.createTransaction)
}

func (h *Handler) createTransaction(w http.ResponseWriter, r *http.Request) {
	var t transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.transactionSvc.Create(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) listTransactions(w http.ResponseWriter, r *http.Request) {

	transactions := h.transactionSvc.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
