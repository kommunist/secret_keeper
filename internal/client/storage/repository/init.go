package repository

import (
	"database/sql"
	"os"
	"secret_keeper/internal/client/logger"

	_ "github.com/lib/pq"
)

type Storage struct {
	driver *sql.DB
}

func Make(dsn string) Storage {
	driver, err := sql.Open("postgres", dsn)

	if err != nil {
		logger.Logger.Error("Error when open database", "err", err)
		os.Exit(1) // TODO сделать вынос ошибки
	}

	return Storage{driver: driver}
}
