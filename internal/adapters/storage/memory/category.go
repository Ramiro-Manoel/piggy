package memory

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/category"
)

type categoryRepository struct {
	categories []category.Category
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{categories: make([]category.Category, 0)}
}

func (r *categoryRepository) Save(c category.Category) error {
	r.categories = append(r.categories, c)
	return nil
}

func (r *categoryRepository) Read(id string) (category.Category, error) {
	for i := range r.categories {
		if r.categories[i].ID == id {
			return r.categories[i], nil
		}
	}
	return category.Category{}, fmt.Errorf("category with id %s not found", id)
}

func (r *categoryRepository) List() []category.Category {
	categories := append([]category.Category{}, r.categories...)
	return categories
}
