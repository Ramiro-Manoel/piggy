package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
	"github.com/jackc/pgx/v5"
)

type transactionRepository struct {
	db *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) scan(row pgx.Row) (transaction.Transaction, error) {
	var t transaction.Transaction
	err := row.Scan(
		&t.ID,
		&t.Ref.ExternalID,
		&t.Ref.Source,
		&t.Description,
		&t.Amount,
		&t.Date,
		&t.CategoryID,
		&t.AccountID)
	if err != nil {
		return transaction.Transaction{}, err
	}
	return t, nil
}

func (r *transactionRepository) Save(t transaction.Transaction) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO transactions (id, external_id, source, description, amount, date, category_id, account_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT (external_id, source) DO NOTHING
		`, t.ID, t.Ref.ExternalID, t.Ref.Source, t.Description, t.Amount, t.Date, t.CategoryID, t.AccountID)

	return err
}

func (r *transactionRepository) Read(id string) (transaction.Transaction, error) {
	row := r.db.QueryRow(context.Background(), `
	SELECT id, external_id, source, description, amount, date, category_id, account_id
	FROM transactions
	WHERE id = $1
	`, id)

	t, err := r.scan(row)
	if err != nil {
		return transaction.Transaction{}, err
	}
	return t, nil
}
func (r *transactionRepository) List() []transaction.Transaction {
	rows, err := r.db.Query(context.Background(), `
	SELECT id, external_id, source, description, amount, date, category_id, account_id
	FROM transactions
	`)
	if err != nil {
		return []transaction.Transaction{}
	}

	var transactions []transaction.Transaction
	for rows.Next() {
		t, err := r.scan(rows)
		if err != nil {
			return []transaction.Transaction{}
		}
		transactions = append(transactions, t)
	}
	return transactions
}
