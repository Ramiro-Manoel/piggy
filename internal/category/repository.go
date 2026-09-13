package category

type repository interface {
	Read(id string) (Category, error)
	Save(c Category) error
	List() []Category
}
