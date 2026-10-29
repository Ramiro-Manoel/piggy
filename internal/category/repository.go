package category

import "errors"

var ErrNotFound = errors.New("category not found")

type repository interface {
	Read(id string) (Category, error)
	Create(Category) error
	List() ([]Category, error)
}
