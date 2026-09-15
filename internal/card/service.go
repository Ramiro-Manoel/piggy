package card

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
	return s.repo.Save(c)
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
		err = s.repo.Save(c)
		if err != nil {
			return err
		}
	}
	return nil
}
