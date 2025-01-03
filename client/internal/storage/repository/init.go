package repository

import (
	"client/internal/logger"
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	driver *sql.DB
}

func Make(dsn string) (Storage, error) {
	driver, err := sql.Open("postgres", dsn)

	if err != nil {
		logger.Logger.Error("Error when open database", "err", err)
		return Storage{}, err
	}

	return Storage{driver: driver}, nil
}
