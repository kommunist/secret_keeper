package repository

import (
	"client/internal/models"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSecretsUpsert(t *testing.T) {
	exList := []struct {
		exName    string
		beginErr  error
		execErr   error
		commitErr error
	}{
		{
			exName: "create_new_simple_happy_path",
		},
		{
			exName:   "create_new_when_begin_err",
			beginErr: errors.New("qqq"),
		},
		{
			exName:  "create_new_stor_return_error",
			execErr: errors.New("some error"),
		},
		{
			exName:    "create_new_stor_commit_return_error",
			commitErr: errors.New("some error"),
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
			secret := models.Secret{
				ID: "123123", Name: "nameOfSecret", Pass: "passOfSecret",
				Meta: "metaOfSecret", UserID: "userID", Version: "version",
			}

			stor := Storage{driver: db}

			mock.ExpectBegin().WillReturnError(ex.beginErr)
			if ex.beginErr != nil {
				err = stor.SecretsUpsert(ctx, []models.Secret{secret})
				assert.Error(t, err)

				return
			}

			exp := mock.ExpectExec(upsertSQL).WithArgs(
				secret.ID, secret.Name, secret.Pass, secret.Meta, secret.UserID, secret.Version,
			)

			if ex.execErr != nil {
				exp.WillReturnError(ex.execErr)
				mock.ExpectRollback()
			} else {
				exp.WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit().WillReturnError(ex.commitErr)
				if ex.commitErr != nil {
					err = stor.SecretsUpsert(ctx, []models.Secret{secret})
					assert.Error(t, err)

					return
				}

			}

			err = stor.SecretsUpsert(ctx, []models.Secret{secret})

			if ex.execErr != nil {
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

			exp := mock.ExpectQuery(listSQL).WithArgs(userID, "0")

			id := "id"
			name := "name"
			pass := "pass"
			meta := "meta"
			version := "version"

			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				rows :=
					sqlmock.NewRows([]string{"id", "name", "pass", "meta", "version"}).
						AddRow(id, name, pass, meta, version)

				exp.WillReturnRows(rows)
			}

			secrets, err := stor.SecretList(ctx, userID, "0")

			if ex.storErr == nil {
				assert.Equal(t, 1, len(secrets))
				assert.NoError(t, err)

				assert.Equal(t, id, secrets[0].ID)
				assert.Equal(t, name, secrets[0].Name)
				assert.Equal(t, pass, secrets[0].Pass)
				assert.Equal(t, meta, secrets[0].Meta)
				assert.Equal(t, version, secrets[0].Version)

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
				assert.Equal(t, version, secret.Version)

			} else {
				assert.Error(t, err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}

}
