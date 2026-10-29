package transaction

type accountRepository interface {
	Create(AccountTransaction) error
	Read(id string) (AccountTransaction, error)
	List(AccountFilter) ([]AccountTransaction, error)
}

type cardRepository interface {
	Create(CardTransaction) error
	Read(id string) (CardTransaction, error)
	List(CardFilter) ([]CardTransaction, error)
}
