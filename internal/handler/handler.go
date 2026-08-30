package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

type Handler struct {
	transactionSvc transactionService
	categorySvc    categoryService
	accountSvc     accountService
}

func NewHandler(transactionSvc transactionService, categorySvc categoryService, accountSvc accountService) *Handler {
	return &Handler{
		transactionSvc: transactionSvc,
		categorySvc:    categorySvc,
		accountSvc:     accountSvc,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /transactions", h.listTransactions)
	mux.HandleFunc("POST /transactions", h.createTransaction)

	mux.HandleFunc("GET /categories", h.listCategories)
	mux.HandleFunc("POST /categories", h.createCategory)

	mux.HandleFunc("GET /accounts", h.listAccounts)
	mux.HandleFunc("POST /accounts", h.createAccount)
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

func (h *Handler) createTransaction(w http.ResponseWriter, r *http.Request) {
	t, err := decode[transaction.Transaction](w, r)
	if err != nil {
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
	writeJSON(w, h.transactionSvc.List())
}

func (h *Handler) createCategory(w http.ResponseWriter, r *http.Request) {
	c, err := decode[category.Category](w, r)
	if err != nil {
		return
	}

	err = h.categorySvc.Create(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) listCategories(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.categorySvc.List())
}

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
