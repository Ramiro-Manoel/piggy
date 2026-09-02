package pluggy

import (
	"math"
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func toTransaction(pt pluggyTransaction) (transaction.Transaction, error) {
	date, err := time.Parse(time.RFC3339Nano, pt.Date)
	if err != nil {
		return transaction.Transaction{}, err
	}

	return transaction.Transaction{
		Ref: external.Reference{
			ExternalID: pt.ID,
			Source:     source},
		Description: pt.Description,
		Date:        date,
		Amount:      int64(math.Round(pt.Amount * 100)),
		AccountID:   pt.AccountID,
	}, nil
}

func toTransactions(pts []pluggyTransaction) ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	for _, pt := range pts {
		t, err := toTransaction(pt)
		if err != nil {
			return []transaction.Transaction{}, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
