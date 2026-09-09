package transaction

type accountRepository interface {
	Save(t AccountTransaction) error
	Read(id string) (AccountTransaction, error)
	List() []AccountTransaction
}

type cardRepository interface {
	Save(t CardTransaction) error
	Read(id string) (CardTransaction, error)
	List() []CardTransaction
}
