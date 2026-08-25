package transaction

type Repository interface {
	Save(t Transaction) error
	Read(id string) (Transaction, error)
	List() []Transaction
}