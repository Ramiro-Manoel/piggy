package transaction

import "fmt"

type CardService struct {
	repo     cardRepository
	provider cardFinanceProvider
	card     cardReader
	invoice  invoiceService
}

func NewCardService(repo cardRepository, provider cardFinanceProvider, card cardReader, invoice invoiceService) *CardService {
	return &CardService{
		repo:     repo,
		provider: provider,
		card:     card,
		invoice:  invoice,
	}
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
	transactions, err := s.provider.FetchCardTransactions(cardID)
	if err != nil {
		return fmt.Errorf("Sync: fetch transactions: %w", err)
	}

	c, err := s.card.Read(cardID)
	if err != nil {
		return fmt.Errorf("Sync: read card %s: %w", cardID, err)
	}

	for _, t := range transactions {
		inv, err := s.invoice.FindOrCreate(c, t.Date)
		if err != nil {
			return fmt.Errorf("Sync: find or create invoice for transaction %s: %w", t.ID, err)
		}
		t.InvoiceID = inv.ID
		if err := s.repo.Save(t); err != nil {
			return fmt.Errorf("Sync: save transaction %s: %w", t.ID, err)
		}
	}
	return nil
}
