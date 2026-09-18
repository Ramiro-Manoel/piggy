package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	accountTransactionSvc accountTransactionService
	cardTransactionSvc    cardTransactionService
	categorySvc           categoryService
	accountSvc            accountService
	institutionID         string
}

func NewHandler(
	accountTransactionSvc accountTransactionService,
	cardTransactionSvc cardTransactionService,
	categorySvc categoryService,
	accountSvc accountService,
	institutionID string,

) *Handler {

	return &Handler{
		accountTransactionSvc: accountTransactionSvc,
		cardTransactionSvc:    cardTransactionSvc,
		categorySvc:           categorySvc,
		accountSvc:            accountSvc,
		institutionID:         institutionID,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /accounts", h.listAccounts)
	mux.HandleFunc("POST /accounts", h.createAccount)
	mux.HandleFunc("POST /accounts/sync", h.syncAccounts)
	mux.HandleFunc("POST /accounts/{accountID}/transactions/sync", h.syncAccountTransactions)

	mux.HandleFunc("GET /transactions", h.listAccountTransactions)
	mux.HandleFunc("POST /transactions", h.createAccountTransaction)

	mux.HandleFunc("GET /categories", h.listCategories)
	mux.HandleFunc("POST /categories", h.createCategory)

}

func decode[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var v T
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return v, err
	}
	return v, nil
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
