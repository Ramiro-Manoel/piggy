package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/jackc/pgx/v5"
)

func toCategory(row categoryRow) category.Category {
	return category.Category{
		ID:       row.ID,
		Name:     row.Name,
		ParentID: row.ParentID,
	}
}

type categoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Save(c category.Category) error {
	_, err := r.db.Exec(context.Background(), `
	INSERT INTO categories(id, name, parent_id)
		VALUES ($1, $2, $3)
	`, c.ID, c.Name, c.ParentID)

	return err
}

func (r *categoryRepository) Read(id string) (category.Category, error) {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM categories WHERE id = $1
	`, id)
	if err != nil {
		return category.Category{}, err
	}
	defer rows.Close()

	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[categoryRow])
	if err != nil {
		return category.Category{}, err
	}
	return toCategory(row), nil
}

func (r *categoryRepository) List() []category.Category {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM categories
	`)
	if err != nil {
		return []category.Category{}
	}
	defer rows.Close()

	categoryRows, err := pgx.CollectRows(rows, pgx.RowToStructByName[categoryRow])
	if err != nil {
		return []category.Category{}
	}

	var categories []category.Category
	for _, row := range categoryRows {
		categories = append(categories, toCategory(row))
	}
	return categories
}
