package pluggy

import (
	"math"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/external"
)

func (c *client) FetchAccounts(itemID string) ([]account.Account, error) {
	accountsResp, err := c.fetchAccounts(itemID, accountTypeBank)
	if err != nil {
		return []account.Account{}, err
	}

	return toAccounts(accountsResp.Results), nil
}

func toAccount(pa pluggyAccount) account.Account {
	return account.Account{
		Ref: external.Reference{
			ExternalID: pa.ID,
			Source:     source},
		Number:  pa.Number,
		Name:    pa.Name,
		Balance: int64(math.Round(pa.Balance * 100)),
		Owner:   pa.Owner,
	}
}

func toAccounts(pas []pluggyAccount) []account.Account {
	var accounts []account.Account
	for _, pa := range pas {
		accounts = append(accounts, toAccount(pa))
	}
	return accounts
}
