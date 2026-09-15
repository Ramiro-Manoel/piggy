package pluggy

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

func (c *client) FetchTransactions(accountID string) ([]transaction.AccountTransaction, error) {
	url := baseURL + transactionsPath + "?accountId=" + accountID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []transaction.AccountTransaction{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return []transaction.AccountTransaction{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return []transaction.AccountTransaction{},
			fmt.Errorf("pluggy fetch trasactions failed: status %d", resp.StatusCode)
	}

	var transactionsResp transactionsResponse
	err = json.NewDecoder(resp.Body).Decode(&transactionsResp)
	if err != nil {
		return []transaction.AccountTransaction{}, err
	}

	return toTransactions(transactionsResp.Results)
}

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
