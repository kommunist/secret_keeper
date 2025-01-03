package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
				assert.NoError(t, err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}

}

func TestSecretList(t *testing.T) {
	exList := []struct {
		exName string

		storErr error
	}{
		{
			exName:  "simple_happy_path",
			storErr: nil,
		},
		{
			exName:  "when_stor_returned_err",
			storErr: errors.New("err"),
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
			userID := "123"

			stor := Storage{driver: db}

			exp := mock.ExpectQuery(listSQL).WithArgs(userID)

			id := "id"
			name := "name"
			pass := "pass"
			meta := "meta"
			version := "version"

			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				exp.WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "pass", "meta", "version"}).
						AddRow(id, name, pass, meta, version),
				)
			}

			secrets, err := stor.SecretList(ctx, userID)

			if ex.storErr == nil {
				assert.Equal(t, 1, len(secrets))
				assert.NoError(t, err)

				assert.Equal(t, id, secrets[0].ID)
				assert.Equal(t, name, secrets[0].Name)
				assert.Equal(t, pass, secrets[0].Pass)
				assert.Equal(t, meta, secrets[0].Meta)
				assert.Equal(t, version, secrets[0].Ver)

			} else {
				assert.Equal(t, 0, len(secrets))
				assert.Error(t, err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}

}

func TestSecretShow(t *testing.T) {
	exList := []struct {
		exName string

		storErr error
	}{
		{
			exName:  "simple_happy_path",
			storErr: nil,
		},
		{
			exName:  "when_stor_returned_err",
			storErr: errors.New("err"),
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
			id := "id"

			stor := Storage{driver: db}

			exp := mock.ExpectQuery(showSQL).WithArgs(id)

			name := "name"
			pass := "pass"
			meta := "meta"
			version := "version"

			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				exp.WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "pass", "meta", "version"}).
						AddRow(id, name, pass, meta, version),
				)
			}

			secret, err := stor.SecretShow(ctx, id)

			if ex.storErr == nil {
				assert.NoError(t, err)

				assert.Equal(t, id, secret.ID)
				assert.Equal(t, name, secret.Name)
				assert.Equal(t, pass, secret.Pass)
				assert.Equal(t, meta, secret.Meta)
				assert.Equal(t, version, secret.Ver)

			} else {
				assert.Error(t, err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}

}
