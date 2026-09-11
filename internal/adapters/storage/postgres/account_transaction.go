package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
	"github.com/jackc/pgx/v5"
)

type accountTransactionRepository struct {
	db *pgx.Conn
}

func NewAccountTransactionRepository(db *pgx.Conn) *accountTransactionRepository {
	return &accountTransactionRepository{db: db}
}

func (r *accountTransactionRepository) scan(row pgx.Row) (transaction.AccountTransaction, error) {
	var t transaction.AccountTransaction
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
		return transaction.AccountTransaction{}, err
	}
	return t, nil
}

func (r *accountTransactionRepository) Save(t transaction.AccountTransaction) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO account_transactions (id, external_id, source, description, amount, date, category_id, account_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT (external_id, source) DO NOTHING
		`, t.ID, t.Ref.ExternalID, t.Ref.Source, t.Description, t.Amount, t.Date, t.CategoryID, t.AccountID)

	return err
}

func (r *accountTransactionRepository) Read(id string) (transaction.AccountTransaction, error) {
	row := r.db.QueryRow(context.Background(), `
	SELECT id, external_id, source, description, amount, date, category_id, account_id
	FROM account_transactions
	WHERE id = $1
	`, id)

	t, err := r.scan(row)
	if err != nil {
		return transaction.AccountTransaction{}, err
	}
	return t, nil
}
func (r *accountTransactionRepository) List() []transaction.AccountTransaction {
	rows, err := r.db.Query(context.Background(), `
	SELECT id, external_id, source, description, amount, date, category_id, account_id
	FROM account_transactions
	`)
	if err != nil {
		return []transaction.AccountTransaction{}
	}
	defer rows.Close()

	var transactions []transaction.AccountTransaction
	for rows.Next() {
		t, err := r.scan(rows)
		if err != nil {
			return []transaction.AccountTransaction{}
		}
		transactions = append(transactions, t)
	}
	return transactions
}
