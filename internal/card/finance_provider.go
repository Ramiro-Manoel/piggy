package card

type financeProvider interface {
	FetchCards() []Card
}
