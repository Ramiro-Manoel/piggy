package handler

import (
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/category"
)

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
