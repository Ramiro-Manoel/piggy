package account

import "fmt"

type service struct {
	repo     repository
	provider financeProvider
}

func NewService(r repository, p financeProvider) *service {
	return &service{
		repo:     r,
		provider: p}
}

func (s *service) Create(a Account) error {
	return s.repo.Save(a)
}

func (s *service) Read(id string) (Account, error) {
	return s.repo.Read(id)
}

func (s *service) List() []Account {
	return s.repo.List()
}

func (s *service) Sync(institutionID string) error {
	accounts, err := s.provider.FetchAccounts(institutionID)
	if err != nil {
		return err
	}

	for _, a := range accounts {
		err = s.repo.Save(a)
		if err != nil {
			return fmt.Errorf("save account %s: %w", a.ID, err)
		}
	}
	return nil
}
