package postgres

import (
	"context"

	"github.com/Ramiro-Manoel/piggy/internal/category"
	"github.com/jackc/pgx/v5"
)

type categoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) scan(row pgx.Row) (category.Category, error) {
	var c category.Category
	err := row.Scan(&c.ID, &c.Name, &c.ParentID)
	if err != nil {
		return category.Category{}, err
	}
	return c, nil
}

func (r *categoryRepository) Save(c category.Category) error {
	_, err := r.db.Exec(context.Background(), `
	INSERT INTO categories(id, name, parent_id)
		VALUES ($1, $2, $3)
	`, c.ID, c.Name, c.ParentID)

	return err
}

func (r *categoryRepository) Read(id string) (category.Category, error) {
	row := r.db.QueryRow(context.Background(), `
	SELECT * FROM categories 
		WHERE id = $1
	`, id)

	c, err := r.scan(row)
	if err != nil {
		return category.Category{}, err
	}

	return c, nil
}

func (r *categoryRepository) List() []category.Category {
	rows, err := r.db.Query(context.Background(), `
	SELECT * FROM categories
	`)
	if err != nil {
		return []category.Category{}
	}

	var categories []category.Category
	for rows.Next() {
		c, err := r.scan(rows)
		if err != nil {
			return []category.Category{}
		}
		categories = append(categories, c)
	}
	return categories
}
