package transaction

type accountFinanceProvider interface {
	FetchTransactions(string) ([]AccountTransaction, error)
}

type cardFinanceProvider interface {
	FetchTransactions(string) ([]CardTransaction, error)
}
