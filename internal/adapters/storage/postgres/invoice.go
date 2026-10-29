package postgres

import (
	"context"
	"errors"

	"github.com/Ramiro-Manoel/piggy/internal/external"
	"github.com/Ramiro-Manoel/piggy/internal/invoice"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func toInvoice(row invoiceRow) invoice.Invoice {
	i := invoice.Invoice{
		ID:        row.ID,
		CardID:    row.CardID,
		CloseDate: row.CloseDate,
		PayDate:   row.PayDate,
		Amount:    row.Amount,
	}
	if row.ExternalID != nil && row.Source != nil {
		i.Ref = external.Reference{ExternalID: *row.ExternalID, Source: *row.Source}
	}
	return i
}

type invoiceRepository struct {
	db *pgxpool.Pool
}

func NewInvoiceRepository(db *pgxpool.Pool) *invoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) Create(i invoice.Invoice) error {
	_, err := r.db.Exec(context.Background(), `
	INSERT INTO invoices(id, external_id, source, card_id, close_date, pay_date, amount)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, i.ID, nullIfEmpty(i.Ref.ExternalID), nullIfEmpty(i.Ref.Source),
		i.CardID, i.CloseDate, i.PayDate, i.Amount)

	return err
}

func (r *invoiceRepository) Update(i invoice.Invoice) error {
	tag, err := r.db.Exec(context.Background(), `
	UPDATE invoices SET
		external_id = $2,
		source      = $3,
		close_date  = $4,
		pay_date    = $5,
		amount      = $6
	WHERE id = $1
	`, i.ID, nullIfEmpty(i.Ref.ExternalID), nullIfEmpty(i.Ref.Source),
		i.CloseDate, i.PayDate, i.Amount)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return invoice.ErrNotFound
	}
	return nil
}

func (r *invoiceRepository) Read(id string) (invoice.Invoice, error) {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM invoices WHERE id = $1
	`, id)
	if err != nil {
		return invoice.Invoice{}, err
	}
	defer rows.Close()

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[invoiceRow])
	if err != nil {
		return invoice.Invoice{}, err
	}
	return toInvoice(row), nil
}

func (r *invoiceRepository) ReadByRef(ref external.Reference) (invoice.Invoice, error) {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM invoices WHERE external_id = $1 AND source = $2
	`, ref.ExternalID, ref.Source)
	if err != nil {
		return invoice.Invoice{}, err
	}
	defer rows.Close()

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[invoiceRow])
	if errors.Is(err, pgx.ErrNoRows) {
		return invoice.Invoice{}, invoice.ErrNotFound
	}
	if err != nil {
		return invoice.Invoice{}, err
	}
	return toInvoice(row), nil
}

func (r *invoiceRepository) List(f invoice.Filter) ([]invoice.Invoice, error) {
	var cardIDs []string
	if len(f.CardIDs) > 0 {
		cardIDs = f.CardIDs
	}

	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM invoices 
	WHERE 	($1::text[] IS NULL OR card_id = ANY($1))
		AND ($2::date   IS NULL OR close_date >= $2)
		AND ($3::date   IS NULL OR close_date <= $3)
	`, cardIDs, f.CloseDate.Start, f.CloseDate.End)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoiceRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[invoiceRow])
	if err != nil {
		return nil, err
	}

	invoices := make([]invoice.Invoice, 0, len(invoiceRows))
	for _, row := range invoiceRows {
		invoices = append(invoices, toInvoice(row))
	}
	return invoices, nil
}

func nullIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
