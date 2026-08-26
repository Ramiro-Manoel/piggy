package memory

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/category"
)

var _ category.Repository = (*CategoryRepository)(nil)

type CategoryRepository struct {
	categories []category.Category
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{categories: make([]category.Category, 0)}
}

func (r *CategoryRepository) Save(c category.Category) error {
	r.categories = append(r.categories, c)
	return nil
}

func (r *CategoryRepository) Read(id string) (category.Category, error) {
	for i := range r.categories {
		if r.categories[i].ID == id {
			return r.categories[i], nil
		}
	}
	return category.Category{}, fmt.Errorf("category with id %s not found", id)
}

func (r *CategoryRepository) List() []category.Category {
	categories := append([]category.Category{}, r.categories...)
	return categories
}
