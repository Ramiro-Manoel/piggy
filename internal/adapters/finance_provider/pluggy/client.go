package pluggy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL          = "https://api.pluggy.ai"
	authPath         = "/auth"
	accountsPath     = "/accounts"
	transactionsPath = "/v2/transactions"
)
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
	url := baseURL + authPath

	body, err := json.Marshal(authRequest{ClientId: c.clientID, ClientSecret: c.clientSecret})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
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

func (c *client) fetchAccounts(itemID string, t accountType) (accountsResponse, error) {
	url := baseURL + accountsPath + "?itemId=" + itemID

	if t != accountTypeAll {
		url += "&type=" + string(t)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return accountsResponse{}, err
	}

	resp, err := c.do(req)
	if err != nil {
		return accountsResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return accountsResponse{},
			fmt.Errorf("pluggy fetch accounts failed: status %d", resp.StatusCode)
	}

	var accountsResp accountsResponse
	err = json.NewDecoder(resp.Body).Decode(&accountsResp)
	if err != nil {
		return accountsResponse{}, err
	}

	return accountsResp, nil
}
