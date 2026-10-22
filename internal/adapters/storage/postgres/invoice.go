package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/invoice"
	"github.com/jackc/pgx/v5"
)

func toInvoice(row invoiceRow) invoice.Invoice {
	return invoice.Invoice{
		ID:        row.ID,
		CardID:    row.CardID,
		CloseDate: row.CloseDate,
		PayDate:   row.PayDate,
		Amount:    row.Amount,
		PaidAt:    row.PaidAt,
	}
}

type invoiceRepository struct {
	db *pgx.Conn
}

func NewInvoiceRepository(db *pgx.Conn) *invoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) Save(i invoice.Invoice) error {
	_, err := r.db.Exec(context.Background(), `
	INSERT INTO invoices(id, card_id, close_date, pay_date, amount, paid_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, i.ID, i.CardID, i.CloseDate, i.PayDate, i.Amount, i.PaidAt)

	return err
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
