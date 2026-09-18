package pluggy

type accountType string

const (
	accountTypeAll    accountType = ""
	accountTypeBank   accountType = "BANK"
	accountTypeCredit accountType = "CREDIT"
)

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
	ID                 string                   `json:"id"`
	Description        string                   `json:"description"`
	Amount             float64                  `json:"amount"`
	Date               string                   `json:"date"`
	AccountID          string                   `json:"accountId"`
	CreditCardMetadata pluggyCreditCardMetadata `json:"creditCardMetadata"`
}

type pluggyCreditCardMetadata struct {
	InstallmentNumber int    `json:"installmentNumber"`
	TotalInstallments int    `json:"totalInstallments"`
	CardNumber        string `json:"cardNumber"`
}

type accountsResponse struct {
	Results []pluggyAccount `json:"results"`
}

type pluggyAccount struct {
	ID         string           `json:"id"`
	Type       string           `json:"type"`
	Number     string           `json:"number"`
	Name       string           `json:"name"`
	Balance    float64          `json:"balance"`
	Owner      string           `json:"owner"`
	CreditData pluggyCreditData `json:"creditData"`
}

type pluggyCreditData struct {
	Brand            string  `json:"brand"`
	CreditLimit      float64 `json:"creditLimit"`
	AvailableLimit   float64 `json:"availableCreditLimit"`
	BalanceCloseDate string  `json:"balanceCloseDate"`
	BalanceDueDate   string  `json:"balanceDueDate"`
}
