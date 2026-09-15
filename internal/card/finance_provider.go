package card

type financeProvider interface {
	FetchCards(string) ([]Card, error)
}
