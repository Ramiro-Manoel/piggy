package transaction

type accountRepository interface {
	Create(AccountTransaction) error
	Read(id string) (AccountTransaction, error)
	ListByAccount(accountID string) ([]AccountTransaction, error)
}

type cardRepository interface {
	Create(CardTransaction) error
	Read(id string) (CardTransaction, error)
	ListByInvoice(invoiceID string) ([]CardTransaction, error)
}
