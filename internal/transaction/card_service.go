package transaction

import "fmt"

type CardService struct {
	repo     cardRepository
	provider cardFinanceProvider
}

func NewCardService(repo cardRepository, provider cardFinanceProvider) *CardService {
	return &CardService{
		repo:     repo,
		provider: provider}
}

func (s *CardService) Create(t CardTransaction) error {
	return s.repo.Save(t)
}

func (s *CardService) List() []CardTransaction {
	return s.repo.List()
}

func (s *CardService) Read(id string) (CardTransaction, error) {
	return s.repo.Read(id)
}

func (s *CardService) Sync(cardID string) error {
	transactions, err := s.provider.FetchTransactions(cardID)
	if err != nil {
		return err
	}

	for _, t := range transactions {
		err = s.repo.Save(t)
		if err != nil {
			return fmt.Errorf("save card transaction %s: %w", t.ID, err)
		}
	}
	return nil
}
