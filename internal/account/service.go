package account

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
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
