package memory

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/account"
)

type AccountRepository struct {
	accounts []account.Account
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{accounts: make([]account.Account, 0)}
}

func (r *AccountRepository) Save(a account.Account) error {
	r.accounts = append(r.accounts, a)
	return nil
}

func (r *AccountRepository) Read(id string) (account.Account, error) {
	for i := range r.accounts {
		if r.accounts[i].ID == id {
			return r.accounts[i], nil
		}
	}
	return account.Account{}, fmt.Errorf("category with id %s not found", id)
}

func (r *AccountRepository) List() []account.Account {
	accounts := append([]account.Account{}, r.accounts...)
	return accounts
}
