package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSecretUpsert(t *testing.T) {
	exList := []struct {
		exName  string
		id      string
		storErr error
	}{
		{
			exName:  "create_new_simple_happy_path",
			id:      "",
			storErr: nil,
		},
		{
			exName:  "create_new_stor_return_error",
			id:      "",
			storErr: errors.New("some error"),
		},
		{
			exName:  "update_new_simple_happy_path",
			id:      "123123",
			storErr: nil,
		},
		{
			exName:  "update_new_stor_return_error",
			id:      "123123",
			storErr: errors.New("some error"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.exName, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.Background()
			name := "nameOfSecret"
			pass := "passOfSecret"
			meta := "metaOfSecret"
			userID := "userUd"
			version := "version"

			stor := Storage{driver: db}

			exp := mock.ExpectExec(upsertSQL)

			if ex.id == "" {
				exp = exp.WithArgs(
					sqlmock.AnyArg(), name, pass, meta, userID, version,
				)
			} else {
				exp = exp.WithArgs(
					ex.id, name, pass, meta, userID, version,
				)
			}

			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				exp.WillReturnResult(sqlmock.NewResult(1, 1))
			}

			err = stor.SecretUpsert(ctx, ex.id, name, pass, meta, userID, version)

			if ex.storErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoErr(t, err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}

}
