package invoice

type repository interface {
	Read(string) (Invoice, error)
	Save(Invoice) error
	List(Filter) ([]Invoice, error)
}
