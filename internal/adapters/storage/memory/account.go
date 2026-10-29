package memory

import "github.com/Ramiro-Manoel/piggy/internal/account"

type accountRepository struct {
	accounts []account.Account
}

func NewAccountRepository() *accountRepository {
	return &accountRepository{accounts: make([]account.Account, 0)}
}

func (r *accountRepository) Create(a account.Account) error {
	r.accounts = append(r.accounts, a)
	return nil
}

func (r *accountRepository) Read(id string) (account.Account, error) {
	for i := range r.accounts {
		if r.accounts[i].ID == id {
			return r.accounts[i], nil
		}
	}
	return account.Account{}, account.ErrNotFound
}

func (r *accountRepository) List() ([]account.Account, error) {
	accounts := append([]account.Account{}, r.accounts...)
	return accounts, nil
}
