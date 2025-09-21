package quote

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SaveQuote(ctx context.Context, bid float64) error {
	query := `
	INSERT INTO quotes (bid, timestamp) VALUES (?, ?)
	`
	_, err := r.db.ExecContext(ctx, query, bid, time.Now())
	if err != nil {
		return err
	}
	return nil
}
