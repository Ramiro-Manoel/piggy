package invoice

import (
	"errors"

	"github.com/Ramiro-Manoel/piggy/internal/external"
)

var ErrNotFound = errors.New("invoice not found")

type repository interface {
	Read(string) (Invoice, error)
	Create(Invoice) error
	Update(Invoice) error
	List(Filter) ([]Invoice, error)
	ReadByRef(external.Reference) (Invoice, error)
}
