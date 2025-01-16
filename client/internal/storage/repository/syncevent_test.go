package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetLastSyncEventVersion(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
	}{
		{
			name: "happy_path",
		},
		{
			name:    "when_no_rows",
			storErr: sql.ErrNoRows,
		},
		{
			name:    "when_another_err",
			storErr: errors.New("123"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			stor := Storage{driver: db}
			exp := mock.ExpectQuery(getLastSyncEventQuery).WithArgs("secret")
			if ex.storErr == nil {
				exp.WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("12345"))
			} else {
				exp.WillReturnError(ex.storErr)
			}

			result, err := stor.GetLastSyncEventVersion(context.Background(), "secret")
			if ex.storErr != nil {
				if ex.storErr == sql.ErrNoRows {
					assert.Equal(t, "0", result)
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
				return
			}

			assert.Equal(t, "12345", result)
			assert.NoError(t, err)
		})
	}
}

func TestSaveSyncEvent(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
	}{
		{
			name: "happy_path",
		},
		{
			name:    "when_another_err",
			storErr: errors.New("123"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			stor := Storage{driver: db}
			exp := mock.ExpectExec(saveSyncEvent).WithArgs("secret", "version").WillReturnResult(sqlmock.NewResult(1, 1))
			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			}

			err = stor.SaveSyncEvent(context.Background(), "secret", "version")

			if ex.storErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

		})
	}
}
