package pluggy

import (
	"math"
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/card"
	"github.com/Ramiro-Manoel/piggy/internal/external"
)

func (c *client) FetchCards(itemID string) ([]card.Card, error) {
	accountsResp, err := c.fetchAccounts(itemID, accountTypeCredit)
	if err != nil {
		return []card.Card{}, err
	}

	return toCards(accountsResp.Results)
}

func toCard(pa pluggyAccount) (card.Card, error) {
	dueDay, err := dayFromDate(pa.CreditData.BalanceDueDate)
	if err != nil {
		return card.Card{}, err
	}
	closeDay, err := dayFromDate(pa.CreditData.BalanceCloseDate)
	if err != nil {
		return card.Card{}, err
	}

	return card.Card{
		Ref: external.Reference{
			ExternalID: pa.ID,
			Source:     source},
		Name:           pa.Name,
		Brand:          pa.CreditData.Brand,
		CreditLimit:    int64(math.Round(pa.CreditData.CreditLimit * 100)),
		AvailableLimit: int64(math.Round(pa.CreditData.AvailableLimit * 100)),
		ClosingDay:     closeDay,
		DueDay:         dueDay,
	}, nil
}

func toCards(pas []pluggyAccount) ([]card.Card, error) {
	var cards []card.Card
	for _, pa := range pas {
		a, err := toCard(pa)
		if err != nil {
			return []card.Card{}, err
		}
		cards = append(cards, a)
	}
	return cards, nil
}

// dayFromDate devolve nil quando a data não veio (null no JSON vira "").
func dayFromDate(date string) (*int, error) {
	if date == "" {
		return nil, nil
	}
	t, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return nil, err
	}
	day := t.Day()
	return &day, nil
}
