package transaction

import (
	"fmt"

	"github.com/google/uuid"
)

type cardService struct {
	repo     cardRepository
	provider cardFinanceProvider
	card     cardReader
	invoice  invoiceService
}

func NewCardService(repo cardRepository, provider cardFinanceProvider, card cardReader, invoice invoiceService) *cardService {
	return &cardService{
		repo:     repo,
		provider: provider,
		card:     card,
		invoice:  invoice,
	}
}

func (s *cardService) Create(t CardTransaction) error {
	t.ID = uuid.NewString()
	return s.repo.Create(t)
}

func (s *cardService) Read(id string) (CardTransaction, error) {
	return s.repo.Read(id)
}

func (s *cardService) ListByInvoice(invoiceID string) ([]CardTransaction, error) {
	return s.repo.ListByInvoice(invoiceID)
}

func (s *cardService) Sync(cardID string) error {
	card, err := s.card.Read(cardID)
	if err != nil {
		return fmt.Errorf("Sync: read card: %w", err)
	}

	transactions, err := s.provider.FetchCardTransactions(card.Ref.ExternalID)
	if err != nil {
		return fmt.Errorf("Sync: fetch transactions: %w", err)
	}

	for _, t := range transactions {
		if t.InvoiceMonth == nil {
			continue
		}
		inv, err := s.invoice.FindOrCreate(card, *t.InvoiceMonth)
		if err != nil {
			return fmt.Errorf("Sync: find or create invoice for transaction %s: %w", t.Ref.ExternalID, err)
		}
		t.InvoiceID = inv.ID
		if err := s.Create(t); err != nil {
			return fmt.Errorf("Sync: save transaction %s: %w", t.Ref.ExternalID, err)
		}
	}
	return nil
}
