package card

type repository interface {
	Save(Card) error
	Read(string) (Card, error)
	List() []Card
}
