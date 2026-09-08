package transaction

type financeProvider interface {
	FetchTransactions(accountID string) ([]AccountTransaction, error)
}
