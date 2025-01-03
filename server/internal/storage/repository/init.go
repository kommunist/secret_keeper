package repository

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type Storage struct {
	driver *sql.DB
}

func Make(dsn string) (Storage, error) {
	driver, err := sql.Open("postgres", dsn)

	if err != nil {
		slog.Error("Error when open database", "err", err)
		return Storage{}, err
	}

	return Storage{driver: driver}, nil
}
