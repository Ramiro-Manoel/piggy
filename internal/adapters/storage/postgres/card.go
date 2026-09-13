package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/card"
	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/jackc/pgx/v5"
)

func toCard(row cardRow) card.Card {
	return card.Card{
		ID: row.ID,
		Ref: external.Reference{
			ExternalID: row.ExternalID,
			Source:     row.Source,
		},
		Name:           row.Name,
		Brand:          row.Brand,
		CreditLimit:    row.CreditLimit,
		AvailableLimit: row.AvailableLimit,
		ClosingDay:     row.ClosingDay,
		DueDay:         row.DueDay,
		AccountID:      row.AccountID,
	}
}

type cardRepository struct {
	db *pgx.Conn
}

func NewCardRepository(db *pgx.Conn) *cardRepository {
	return &cardRepository{db: db}
}

func (r *cardRepository) Save(c card.Card) error {
	_, err := r.db.Exec(context.Background(), `
	INSERT INTO cards(
		id, external_id, source, name, brand,
		credit_limit, available_limit, closing_day, due_day, account_id
	)
	VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10)
	ON CONFLICT (external_id, source) DO NOTHING`,
		c.ID, c.Ref.ExternalID, c.Ref.Source, c.Name, c.Brand,
		c.CreditLimit, c.AvailableLimit, c.ClosingDay, c.DueDay, c.AccountID)

	return err
}

func (r *cardRepository) Read(id string) (card.Card, error) {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM cards WHERE id = $1
	`, id)
	if err != nil {
		return card.Card{}, err
	}
	defer rows.Close()

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[cardRow])
	if err != nil {
		return card.Card{}, err
	}
	return toCard(row), nil
}

func (r *cardRepository) List() []card.Card {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM cards
	`)
	if err != nil {
		return []card.Card{}
	}
	defer rows.Close()

	cardRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[cardRow])
	if err != nil {
		return []card.Card{}
	}

	var cards []card.Card
	for _, row := range cardRows {
		cards = append(cards, toCard(row))
	}
	return cards
}
