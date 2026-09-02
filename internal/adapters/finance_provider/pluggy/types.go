package pluggy

type authRequest struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type authResponse struct {
	ApiKey string `json:"apiKey"`
}

type transactionsResponse struct {
	Results []pluggyTransaction `json:"results"`
}

type pluggyTransaction struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
	AccountID   string  `json:"accountId"`
}

type accountsResponse struct {
	Results []pluggyAccount `json:"results"`
}

type pluggyAccount struct {
	ID     string  `json:"id"`
	Number string  `json:"number"`
	Name   string  `json:"name"`
	Blance float64 `json:"balance"`
	Owner  string  `json:"owner"`
}
