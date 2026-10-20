package transaction

import (
	"fmt"

	"github.com/google/uuid"
)

type AccountService struct {
	repo     accountRepository
	provider accountFinanceProvider
}

func NewAccountService(repo accountRepository, provider accountFinanceProvider) *AccountService {
	return &AccountService{
		repo:     repo,
		provider: provider}
}

func (s *AccountService) Create(t AccountTransaction) error {
	t.ID = uuid.NewString()
	return s.repo.Save(t)
}

func (s *AccountService) List() []AccountTransaction {
	return s.repo.List()
}

func (s *AccountService) Read(id string) (AccountTransaction, error) {
	return s.repo.Read(id)
}

func (s *AccountService) Sync(accountID string) error {
	transactions, err := s.provider.FetchAccountTransactions(accountID)
	if err != nil {
		return err
	}

	for _, t := range transactions {
		if err := s.Create(t); err != nil {
			return fmt.Errorf("save account transaction %s: %w", t.Ref.ExternalID, err)
		}
	}
	return nil
}
