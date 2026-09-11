package postgres

import (
	"context"
	"time"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
	"github.com/jackc/pgx/v5"
)

type cardTransactionRow struct {
	ID                string    `db:"id"`
	ExternalID        string    `db:"external_id"`
	Source            string    `db:"source"`
	Description       string    `db:"description"`
	Amount            int64     `db:"amount"`
	Date              time.Time `db:"date"`
	CategoryID        *string   `db:"category_id"`
	InvoiceID         string    `db:"invoice_id"`
	InstallmentNumber int       `db:"installment_number"`
	TotalInstallments int       `db:"total_installments"`
}

func toCardTransaction(row cardTransactionRow) transaction.CardTransaction {
	t := transaction.CardTransaction{
		InvoiceID:         row.InvoiceID,
		InstallmentNumber: row.InstallmentNumber,
		TotalInstallments: row.TotalInstallments,
	}
	t.ID = row.ID
	t.Ref.ExternalID = row.ExternalID
	t.Ref.Source = row.Source
	t.Description = row.Description
	t.Amount = row.Amount
	t.Date = row.Date
	t.CategoryID = row.CategoryID

	return t
}

type cardTransactionRepository struct {
	db *pgx.Conn
}

func NewCardTransactionRepository(db *pgx.Conn) *cardTransactionRepository {
	return &cardTransactionRepository{db: db}
}

func (r *cardTransactionRepository) Save(t transaction.CardTransaction) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO card_transactions (
			id, external_id, source, description, amount,
			date, category_id, invoice_id, installment_number, total_installments
		) VALUES (
		 	$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10)
			ON CONFLICT (external_id, source) DO NOTHING
		`,
		t.ID, t.Ref.ExternalID, t.Ref.Source, t.Description, t.Amount,
		t.Date, t.CategoryID, t.InvoiceID, t.InstallmentNumber, t.TotalInstallments)

	return err
}

func (r *cardTransactionRepository) Read(id string) (transaction.CardTransaction, error) {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM card_transactions WHERE id = $1
	`, id)
	if err != nil {
		return transaction.CardTransaction{}, err
	}

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[cardTransactionRow])
	if err != nil {
		return transaction.CardTransaction{}, err
	}

	return toCardTransaction(row), nil
}

func (r *cardTransactionRepository) List() []transaction.CardTransaction {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM card_transactions
	`)
	if err != nil {
		return []transaction.CardTransaction{}
	}
	defer rows.Close()

	cardTransactionRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[cardTransactionRow])
	if err != nil {
		return []transaction.CardTransaction{}
	}
	var transactions []transaction.CardTransaction
	for _, row := range cardTransactionRows {
		transactions = append(transactions, toCardTransaction(row))
	}
	return transactions
}
