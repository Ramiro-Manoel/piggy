package category

type service struct {
	repo repository
}

func NewService(r repository) *service {
	return &service{repo: r}
}

func (s *service) Create(c Category) error {
	return s.repo.Save(c)
}

func (s *service) Read(id string) (Category, error) {
	return s.repo.Read(id)
}

func (s *service) List() []Category {
	return s.repo.List()
}
