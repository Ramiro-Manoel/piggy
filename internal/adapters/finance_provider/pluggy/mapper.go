package pluggy

import (
	"math"
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func toTransaction(pt pluggyTransaction) (transaction.AccountTransaction, error) {
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

func toTransactions(pts []pluggyTransaction) ([]transaction.AccountTransaction, error) {
	var transactions []transaction.AccountTransaction
	for _, pt := range pts {
		t, err := toTransaction(pt)
		if err != nil {
			return []transaction.AccountTransaction{}, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func toAccount(pa pluggyAccount) (account.Account, error) {
	return account.Account{
		Ref: external.Reference{
			ExternalID: pa.ID,
			Source:     source},
		Number:  pa.Number,
		Name:    pa.Name,
		Balance: int64(math.Round(pa.Balance * 100)),
		Owner:   pa.Owner,
	}, nil
}

func toAccounts(pas []pluggyAccount) ([]account.Account, error) {
	var accounts []account.Account
	for _, pa := range pas {
		a, err := toAccount(pa)
		if err != nil {
			return []account.Account{}, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}
