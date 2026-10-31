package memory

import "github.com/Ramiro-Manoel/piggy/internal/card"

type cardRepository struct {
	cards []card.Card
}

func NewCardRepository() *cardRepository {
	return &cardRepository{cards: make([]card.Card, 0)}
}

func (r *cardRepository) Save(c card.Card) error {
	for i := range r.cards {
		if r.cards[i].Ref == c.Ref {
			c.ID = r.cards[i].ID
			if c.ClosingDay == nil {
				c.ClosingDay = r.cards[i].ClosingDay
			}
			if c.DueDay == nil {
				c.DueDay = r.cards[i].DueDay
			}
			if c.AccountID == nil {
				c.AccountID = r.cards[i].AccountID
			}
			r.cards[i] = c
			return nil
		}
	}
	r.cards = append(r.cards, c)
	return nil
}

func (r *cardRepository) Read(id string) (card.Card, error) {
	for i := range r.cards {
		if r.cards[i].ID == id {
			return r.cards[i], nil
		}
	}
	return card.Card{}, card.ErrNotFound
}

func (r *cardRepository) List() ([]card.Card, error) {
	cards := append([]card.Card{}, r.cards...)
	return cards, nil
}
