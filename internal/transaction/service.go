package transaction

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service{
	return &Service{repo: repo}
} 

func (s *Service) Save(t Transaction) error {
	return s.repo.Save(t)
}

func (s *Service) List() []Transaction {
	return s.repo.List()
}