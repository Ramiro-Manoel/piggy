package pluggy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.pluggy.ai"

type authRequest struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type authResponse struct {
	ApiKey string `json:"apiKey"`
}
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
	body, err := json.Marshal(authRequest{ClientId: c.clientID, ClientSecret: c.clientSecret})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", baseURL+"/auth", bytes.NewReader(body))
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
