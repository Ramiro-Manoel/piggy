package account

import "errors"

var ErrNotFound = errors.New("account not found")

type repository interface {
	Read(id string) (Account, error)
	Create(Account) error
	List() []Account
}
