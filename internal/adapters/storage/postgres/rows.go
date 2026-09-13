package postgres

import "time"

type accountRow struct {
	ID         string `db:"id"`
	ExternalID string `db:"external_id"`
	Source     string `db:"source"`
	Name       string `db:"name"`
	Number     string `db:"number"`
	Owner      string `db:"owner"`
	Balance    int64  `db:"balance"`
}

type cardRow struct {
	ID             string  `db:"id"`
	ExternalID     string  `db:"external_id"`
	Source         string  `db:"source"`
	Name           string  `db:"name"`
	Brand          string  `db:"brand"`
	CreditLimit    int64   `db:"credit_limit"`
	AvailableLimit int64   `db:"available_limit"`
	ClosingDay     int     `db:"closing_day"`
	DueDay         int     `db:"due_day"`
	AccountID      *string `db:"account_id"`
}

type categoryRow struct {
	ID       string  `db:"id"`
	Name     string  `db:"name"`
	ParentID *string `db:"parent_id"`
}

type transactionRow struct {
	ID          string    `db:"id"`
	ExternalID  string    `db:"external_id"`
	Source      string    `db:"source"`
	Description string    `db:"description"`
	Amount      int64     `db:"amount"`
	Date        time.Time `db:"date"`
	CategoryID  *string   `db:"category_id"`
}

type accountTransactionRow struct {
	transactionRow
	AccountID string `db:"account_id"`
}

type cardTransactionRow struct {
	transactionRow
	InvoiceID         string `db:"invoice_id"`
	InstallmentNumber int    `db:"installment_number"`
	TotalInstallments int    `db:"total_installments"`
}
