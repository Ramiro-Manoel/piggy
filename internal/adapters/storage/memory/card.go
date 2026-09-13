package memory

import (
	"fmt"

	"github.com/Ramiro-Manoel/piggy/internal/card"
)

type cardRepository struct {
	cards []card.Card
}

func NewCardRepository() *cardRepository {
	return &cardRepository{cards: make([]card.Card, 0)}
}

func (r *cardRepository) Save(c card.Card) error {
	r.cards = append(r.cards, c)
	return nil
}

func (r *cardRepository) Read(id string) (card.Card, error) {
	for i := range r.cards {
		if r.cards[i].ID == id {
			return r.cards[i], nil
		}
	}
	return card.Card{}, fmt.Errorf("card with id %s not found", id)
}

func (r *cardRepository) List() []card.Card {
	cards := append([]card.Card{}, r.cards...)
	return cards
}
