package transaction

import "fmt"

type Service struct {
	repo     Repository
	provider financeProvider
}

func NewService(repo Repository, provider financeProvider) *Service {
	return &Service{
		repo:     repo,
		provider: provider}
}

func (s *Service) Create(t Transaction) error {
	return s.repo.Save(t)
}

func (s *Service) List() []Transaction {
	return s.repo.List()
}

func (s *Service) Read(id string) (Transaction, error) {
	return s.repo.Read(id)
}

func (s *Service) Sync(accountID string) error {
	transactions, err := s.provider.FetchTransactions(accountID)
	if err != nil {
		return err
	}

	for _, t := range transactions {
		err = s.repo.Save(t)
		if err != nil {
			return fmt.Errorf("save transaction %s: %w", t.ID, err)
		}
	}
	return nil
}
