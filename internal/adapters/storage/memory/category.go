package memory

import "github.com/Ramiro-Manoel/piggy/internal/category"

type categoryRepository struct {
	categories []category.Category
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{categories: make([]category.Category, 0)}
}

func (r *categoryRepository) Create(c category.Category) error {
	r.categories = append(r.categories, c)
	return nil
}

func (r *categoryRepository) Read(id string) (category.Category, error) {
	for i := range r.categories {
		if r.categories[i].ID == id {
			return r.categories[i], nil
		}
	}
	return category.Category{}, category.ErrNotFound
}

func (r *categoryRepository) List() ([]category.Category, error) {
	categories := append([]category.Category{}, r.categories...)
	return categories, nil
}
