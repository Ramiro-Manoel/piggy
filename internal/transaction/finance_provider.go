package transaction

type accountFinanceProvider interface {
	FetchAccountTransactions(string) ([]AccountTransaction, error)
}

type cardFinanceProvider interface {
	FetchCardTransactions(string) ([]CardTransaction, error)
}
