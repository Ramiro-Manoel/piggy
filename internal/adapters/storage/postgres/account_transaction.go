package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/transaction"
	"github.com/jackc/pgx/v5"
)

func toAccountTransaction(row accountTransactionRow) transaction.AccountTransaction {
	t := transaction.AccountTransaction{AccountID: row.AccountID}
	t.ID = row.ID
	t.Ref.ExternalID = row.ExternalID
	t.Ref.Source = row.Source
	t.Description = row.Description
	t.Amount = row.Amount
	t.Date = row.Date
	t.CategoryID = row.CategoryID

	return t
}

type accountTransactionRepository struct {
	db *pgx.Conn
}

func NewAccountTransactionRepository(db *pgx.Conn) *accountTransactionRepository {
	return &accountTransactionRepository{db: db}
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
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM account_transactions WHERE id = $1
	`, id)
	if err != nil {
		return transaction.AccountTransaction{}, err
	}
	defer rows.Close()

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[accountTransactionRow])
	if err != nil {
		return transaction.AccountTransaction{}, err
	}

	return toAccountTransaction(row), nil
}
func (r *accountTransactionRepository) List() []transaction.AccountTransaction {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM account_transactions
	`)
	if err != nil {
		return []transaction.AccountTransaction{}
	}
	defer rows.Close()

	transactionRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[accountTransactionRow])
	if err != nil {
		return []transaction.AccountTransaction{}
	}

	var transactions []transaction.AccountTransaction
	for _, row := range transactionRows {
		transactions = append(transactions, toAccountTransaction(row))
	}
	return transactions
}
