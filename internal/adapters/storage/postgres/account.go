package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/account"
	"github.com/jackc/pgx/v5"
)

type accountRepository struct {
	db *pgx.Conn
}

func NewAccountRepository(db *pgx.Conn) *accountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Save(a account.Account) error {
	_, err := r.db.Exec(context.Background(), `
		INSERT INTO accounts (id, name, number, owner, balance)
		VALUES ($1, $2, $3, $4, $5)
	`, a.ID, a.Name, a.Number, a.Owner, a.Balance)

	return err
}

func (r *accountRepository) Read(id string) (account.Account, error) {
	var a account.Account

	row := r.db.QueryRow(context.Background(), `
	SELECT * FROM accounts 
		WHERE id = $1
	`, id)

	err := row.Scan(&a.ID, &a.Name, &a.Number, &a.Owner, &a.Balance)
	if err != nil {
		return account.Account{}, err
	}
	return a, nil
}

func (r *accountRepository) List() []account.Account {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM accounts 
	`)
	if err != nil {
		return []account.Account{}
	}
	defer rows.Close()

	var accounts []account.Account
	for rows.Next() {
		var a account.Account
		err = rows.Scan(&a.ID, &a.Name, &a.Number, &a.Owner, &a.Balance)
		if err != nil {
			return []account.Account{}
		}
		accounts = append(accounts, a)
	}
	return accounts
}
