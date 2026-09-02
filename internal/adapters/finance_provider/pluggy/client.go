package pluggy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
)

const baseURL = "https://api.pluggy.ai"
const source = "pluggy"

type client struct {
	clientID     string
	clientSecret string
	apiKey       string
}

func NewClient(clientID, clientSecret string) *client {
	return &client{
		clientID:     clientID,
		clientSecret: clientSecret}
}

func (c *client) Authenticate() error {
	url := baseURL + "/auth"

	body, err := json.Marshal(authRequest{ClientId: c.clientID, ClientSecret: c.clientSecret})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("pluggy auth failed: status %d", resp.StatusCode)
	}
	var authResp authResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return err
	}
	c.apiKey = authResp.ApiKey
	return nil
}

func (c *client) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-KEY", c.apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		resp.Body.Close()
		err = c.Authenticate()
		if err != nil {
			return nil, err
		}
		req.Header.Set("X-API-KEY", c.apiKey)
		return http.DefaultClient.Do(req)
	}
	return resp, nil
}

func (c *client) FetchTransactions(accountID string) ([]transaction.Transaction, error) {
	url := baseURL + "/v2/transactions?accountId=" + accountID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []transaction.Transaction{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return []transaction.Transaction{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return []transaction.Transaction{},
			fmt.Errorf("pluggy fetch trasactions failed: status %d", resp.StatusCode)
	}

	var transactionsResp transactionsResponse
	err = json.NewDecoder(resp.Body).Decode(&transactionsResp)
	if err != nil {
		return []transaction.Transaction{}, err
	}
	transactions, err := toTransactions(transactionsResp.Results)
	if err != nil {
		return []transaction.Transaction{}, err
	}

	return transactions, nil
}
