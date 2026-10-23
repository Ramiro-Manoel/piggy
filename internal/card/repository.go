package card

type repository interface {
	Create(Card) error
	Read(string) (Card, error)
	List() []Card
}
