package card

import "errors"

var ErrNotFound = errors.New("card not found")

type repository interface {
	Create(Card) error
	Read(string) (Card, error)
	List() ([]Card, error)
}
