package repository

import (
	"context"
	"errors"
	"server/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSecretUpsert(t *testing.T) {
	exList := []struct {
		exName    string
		id        string
		beginErr  error
		queryErr  error
		commitErr error
	}{
		{
			exName:    "create_new_simple_happy_path",
			id:        "",
			beginErr:  nil,
			queryErr:  nil,
			commitErr: nil,
		},
		{
			exName:    "create_new_begin_err",
			id:        "",
			beginErr:  errors.New("qq"),
			queryErr:  nil,
			commitErr: nil,
		},
		{
			exName:    "create_new_query_err",
			id:        "",
			beginErr:  nil,
			queryErr:  errors.New("qq"),
			commitErr: nil,
		},
		{
			exName:    "create_new_commit_err",
			id:        "",
			beginErr:  nil,
			queryErr:  nil,
			commitErr: errors.New("qq"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.exName, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			list := []models.Secret{
				{
					ID:      "idOfSecret",
					Name:    "nameOfSecret",
					Pass:    "passOfSecret",
					Meta:    "metaOfSecret",
					UserID:  "userUd",
					Version: "version",
				},
			}

			ctx := context.Background()

			stor := Storage{driver: db}

			expBegin := mock.ExpectBegin()
			if ex.beginErr != nil {
				expBegin.WillReturnError(ex.beginErr)
			}

			if ex.beginErr == nil {
				expQuery := mock.ExpectExec(upsertSQL).WithArgs(
					list[0].ID, list[0].Name, list[0].Pass, list[0].Meta, list[0].UserID, list[0].Version,
				)

				if ex.queryErr == nil {
					expQuery.WillReturnResult(sqlmock.NewResult(1, 1))

					expCommit := mock.ExpectCommit()
					if ex.commitErr != nil {
						expCommit.WillReturnError(ex.commitErr)
					}
				} else {
					expQuery.WillReturnError(ex.queryErr)
					mock.ExpectRollback() // не рассматриваю пока вариант, когда rollback упал
				}
			}

			err = stor.SecretUpsert(ctx, list)

			if ex.beginErr == nil && ex.queryErr == nil && ex.commitErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}

}

func TestSecretGet(t *testing.T) {
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
			moreVersion := "5678981234"

			stor := Storage{driver: db}

			exp := mock.ExpectQuery(getSQL).WithArgs(userID, moreVersion)

			id := "id"
			name := "name"
			pass := "pass"
			meta := "meta"
			version := "version"

			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				exp.WillReturnRows(
					sqlmock.NewRows([]string{"id", "name", "pass", "meta", "user_id", "version"}).
						AddRow(id, name, pass, meta, userID, version),
				)
			}

			secrets, err := stor.SecretGet(ctx, userID, moreVersion)

			if ex.storErr == nil {
				assert.Equal(t, 1, len(secrets))
				assert.NoError(t, err)

				assert.Equal(t, id, secrets[0].ID)
				assert.Equal(t, name, secrets[0].Name)
				assert.Equal(t, pass, secrets[0].Pass)
				assert.Equal(t, meta, secrets[0].Meta)
				assert.Equal(t, userID, secrets[0].UserID)
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
