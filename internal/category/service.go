package category

import "github.com/google/uuid"

type service struct {
	repo repository
}

func NewService(r repository) *service {
	return &service{repo: r}
}

func (s *service) Create(c Category) error {
	c.ID = uuid.NewString()
	return s.repo.Create(c)
}

func (s *service) Read(id string) (Category, error) {
	return s.repo.Read(id)
}

func (s *service) List() ([]Category, error) {
	return s.repo.List()
}
