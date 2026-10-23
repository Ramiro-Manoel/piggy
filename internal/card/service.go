package card

import "github.com/google/uuid"

type service struct {
	repo     repository
	provider financeProvider
}

func NewService(r repository, p financeProvider) *service {
	return &service{
		repo:     r,
		provider: p,
	}
}

func (s *service) Create(c Card) error {
	c.ID = uuid.NewString()
	return s.repo.Create(c)
}

func (s *service) Read(id string) (Card, error) {
	return s.repo.Read(id)
}

func (s *service) List() []Card {
	return s.repo.List()
}

func (s *service) Sync(institutionID string) error {
	cards, err := s.provider.FetchCards(institutionID)
	if err != nil {
		return err
	}

	for _, c := range cards {
		if err := s.Create(c); err != nil {
			return err
		}
	}
	return nil
}
