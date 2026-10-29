package transaction

import (
	"fmt"

	"github.com/google/uuid"
)

type accountService struct {
	repo     accountRepository
	provider accountFinanceProvider
	account  accountReader
}

func NewAccountService(repo accountRepository, provider accountFinanceProvider, account accountReader) *accountService {
	return &accountService{
		repo:     repo,
		provider: provider,
		account:  account}
}

func (s *accountService) Create(t AccountTransaction) error {
	t.ID = uuid.NewString()
	return s.repo.Create(t)
}

func (s *accountService) Read(id string) (AccountTransaction, error) {
	return s.repo.Read(id)
}

func (s *accountService) List(f AccountFilter) ([]AccountTransaction, error) {
	return s.repo.List(f)
}

func (s *accountService) Sync(accountID string) error {
	account, err := s.account.Read(accountID)
	if err != nil {
		return fmt.Errorf("Sync: read account: %w", err)
	}

	transactions, err := s.provider.FetchAccountTransactions(account.Ref.ExternalID)
	if err != nil {
		return err
	}

	for _, t := range transactions {
		t.AccountID = account.ID
		if err := s.Create(t); err != nil {
			return fmt.Errorf("save account transaction %s: %w", t.Ref.ExternalID, err)
		}
	}
	return nil
}
