package account

import "fmt"

type Service struct {
	repo     Repository
	provider financeProvider
}

func NewService(r Repository, p financeProvider) *Service {
	return &Service{
		repo:     r,
		provider: p}
}

func (s *Service) Create(a Account) error {
	return s.repo.Save(a)
}

func (s *Service) Read(id string) (Account, error) {
	return s.repo.Read(id)
}

func (s *Service) List() []Account {
	return s.repo.List()
}

func (s *Service) Sync(institutionID string) error {
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
