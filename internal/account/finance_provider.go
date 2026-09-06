package account

type financeProvider interface {
	FetchAccounts(string) ([]Account, error)
}
