package postgres

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	logger *slog.Logger
	db *sqlx.DB
}

func NewPostgresDB(authLink string, logger *slog.Logger) (*Repository, error) {
	db, err := sqlx.Open("postgres", authLink)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
		logger: logger,
	}, nil
}