package transaction

type financeProvider interface {
	FetchTransactions(accountID string) ([]Transaction, error)
}
