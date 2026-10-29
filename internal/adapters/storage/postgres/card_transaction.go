package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

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
	db *pgxpool.Pool
}

func NewCardTransactionRepository(db *pgxpool.Pool) *cardTransactionRepository {
	return &cardTransactionRepository{db: db}
}

func (r *cardTransactionRepository) Create(t transaction.CardTransaction) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO card_transactions (
			id, external_id, source, description, amount,
			date, category_id, invoice_id, installment_number, total_installments
		) VALUES (
		 	$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10)
			ON CONFLICT (external_id, source) WHERE external_id IS NOT NULL DO NOTHING
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
	defer rows.Close()

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[cardTransactionRow])
	if err != nil {
		return transaction.CardTransaction{}, err
	}

	return toCardTransaction(row), nil
}

func (r *cardTransactionRepository) List(f transaction.CardFilter) ([]transaction.CardTransaction, error) {
	var invoiceIDs []string
	if len(f.InvoiceIDs) > 0 {
		invoiceIDs = f.InvoiceIDs
	}

	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM card_transactions
	WHERE ($1::text[] IS NULL OR invoice_id = ANY($1))
	ORDER BY date
	`, invoiceIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactionRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[cardTransactionRow])
	if err != nil {
		return nil, err
	}

	transactions := make([]transaction.CardTransaction, 0, len(transactionRows))
	for _, row := range transactionRows {
		transactions = append(transactions, toCardTransaction(row))
	}
	return transactions, nil
}
