package transaction

type Repository interface {
	Save(t AccountTransaction) error
	Read(id string) (AccountTransaction, error)
	List() []AccountTransaction
}