package category

type Repository interface{
	Read(id string) (Category, error)
	Save(c Category) error
	List() []Category
}