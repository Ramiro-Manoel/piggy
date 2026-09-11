package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/accountTrnansaction"
	"github.com/Ramiro-Manoel/piggy/internal/transaction"
	"github.com/jackc/pgx/v5"
)

type accountTrnansactionRepository struct {
	db *pgx.Conn
}

func NewAccountaccountTrnansactionRepository(db *pgx.Conn) *transactionRepository {
	return &accountTrnansactionRepository{db: db}
}

func (r *accountTrnansactionRepository) scan(row pgx.Row) (transaction.AccountTransaction, error) {
	var t accountTrnansaction.AccountTransaction
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
		return accountTrnansaction.AccountTransaction{}, err
	}
	return t, nil
}

func (r *accountTrnansactionRepository) Save(t transaction.AccountTransaction) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO accountTrnansactions (id, external_id, source, description, amount, date, category_id, account_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			ON CONFLICT (external_id, source) DO NOTHING
		`, t.ID, t.Ref.ExternalID, t.Ref.Source, t.Description, t.Amount, t.Date, t.CategoryID, t.AccountID)

	return err
}

func (r *accountTrnansactionRepository) Read(id string) (transaction.AccountTransaction, error) {
	row := r.db.QueryRow(context.Background(), `
	SELECT id, external_id, source, description, amount, date, category_id, account_id
	FROM accountTrnansactions
	WHERE id = $1
	`, id)

	t, err := r.scan(row)
	if err != nil {
		return accountTrnansaction.AccountTransaction{}, err
	}
	return t, nil
}
func (r *accountTrnansactionRepository) List() []transaction.AccountTransaction {
	rows, err := r.db.Query(context.Background(), `
	SELECT id, external_id, source, description, amount, date, category_id, account_id
	FROM accountTrnansactions
	`)
	if err != nil {
		return []accountTrnansaction.AccountTransaction{}
	}
	defer rows.Close()

	var accountTrnansactions []transaction.AccountTransaction
	for rows.Next() {
		t, err := r.scan(rows)
		if err != nil {
			return []accountTrnansaction.AccountTransaction{}
		}
		accountTrnansactions = append(transactions, t)
	}
	return accountTrnansactions
}
