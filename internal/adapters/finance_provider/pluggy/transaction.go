package pluggy

import (
	"math"
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func (c *client) FetchAccountTransactions(accountID string) ([]transaction.AccountTransaction, error) {
	transactionsResp, err := c.fetchTransactions(accountID)
	if err != nil {
		return []transaction.AccountTransaction{}, err
	}

	return toAccountTransactions(transactionsResp.Results)
}

func toAccountTransaction(pt pluggyTransaction) (transaction.AccountTransaction, error) {
	date, err := time.Parse(time.RFC3339Nano, pt.Date)
	if err != nil {
		return transaction.AccountTransaction{}, err
	}

	t := transaction.AccountTransaction{AccountID: pt.AccountID}
	t.Ref = external.Reference{
		ExternalID: pt.ID,
		Source:     source}
	t.Description = pt.Description
	t.Date = date
	t.Amount = int64(math.Round(pt.Amount * 100))

	return t, nil
}

func toAccountTransactions(pts []pluggyTransaction) ([]transaction.AccountTransaction, error) {
	var transactions []transaction.AccountTransaction
	for _, pt := range pts {
		t, err := toAccountTransaction(pt)
		if err != nil {
			return []transaction.AccountTransaction{}, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (c *client) FetchCardTransactions(accountID string) ([]transaction.CardTransaction, error) {
	transactionsResp, err := c.fetchTransactions(accountID)
	if err != nil {
		return []transaction.CardTransaction{}, err
	}

	return toCardTransactions(transactionsResp.Results)
}

func toCardTransaction(pt pluggyTransaction) (transaction.CardTransaction, error) {
	date, err := time.Parse(time.RFC3339Nano, pt.Date)
	if err != nil {
		return transaction.CardTransaction{}, err
	}

	t := transaction.CardTransaction{
		InstallmentNumber: pt.CreditCardMetadata.InstallmentNumber,
		TotalInstallments: pt.CreditCardMetadata.TotalInstallments,
	}

	t.Ref = external.Reference{
		ExternalID: pt.ID,
		Source:     source}
	t.Description = pt.Description
	t.Date = date
	t.Amount = int64(math.Round(pt.Amount * 100))

	return t, nil
}

func toCardTransactions(pts []pluggyTransaction) ([]transaction.CardTransaction, error) {
	var transactions []transaction.CardTransaction
	for _, pt := range pts {
		t, err := toCardTransaction(pt)
		if err != nil {
			return []transaction.CardTransaction{}, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
