package category

type repository interface {
	Read(id string) (Category, error)
	Create(Category) error
	List() []Category
}
