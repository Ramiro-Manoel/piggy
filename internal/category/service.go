package category

type Service struct{
	repo Repository
}

func NewService(r Repository) *Service{
	return &Service{repo: r}
}

func (s *Service) Create(c Category) error{
	return s.repo.Save(c)
}

func (s *Service) List() []Category{
	return s.repo.List()
}
