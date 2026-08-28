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

func (h *Handler) createCategory(w http.ResponseWriter, r *http.Request) {
	var c category.Category
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
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

	categories := h.categorySvc.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	var a account.Account
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
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

	accounts := h.accountSvc.List()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
