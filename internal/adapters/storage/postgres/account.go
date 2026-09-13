package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/jackc/pgx/v5"
)

func toAccount(row accountRow) account.Account {
	return account.Account{
		ID: row.ID,
		Ref: external.Reference{
			ExternalID: row.ExternalID,
			Source:     row.Source,
		},
		Name:    row.Name,
		Number:  row.Number,
		Owner:   row.Owner,
		Balance: row.Balance,
	}
}

type accountRepository struct {
	db *pgx.Conn
}

func NewAccountRepository(db *pgx.Conn) *accountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Save(a account.Account) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO accounts (id, external_id, source, name, number, owner, balance)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (external_id, source) DO NOTHING
	`, a.ID, a.Ref.ExternalID, a.Ref.Source, a.Name, a.Number, a.Owner, a.Balance)

	return err
}

func (r *accountRepository) Read(id string) (account.Account, error) {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM accounts WHERE id = $1
	`, id)
	if err != nil {
		return account.Account{}, err
	}
	defer rows.Close()

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[accountRow])
	if err != nil {
		return account.Account{}, err
	}
	return toAccount(row), nil
}

func (r *accountRepository) List() []account.Account {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM accounts 
	`)
	if err != nil {
		return []account.Account{}
	}
	defer rows.Close()

	accountRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[accountRow])
	if err != nil {
		return []account.Account{}
	}

	var accounts []account.Account
	for _, row := range accountRows {
		accounts = append(accounts, toAccount(row))
	}
	return accounts
}
