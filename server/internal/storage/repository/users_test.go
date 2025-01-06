package repository

import (
	"context"
	"errors"
	"server/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserSet(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
	}{
		{
			name:    "simple_happy_path",
			storErr: nil,
		},
		{
			name:    "when_stor_return_err",
			storErr: errors.New("some_err"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			user := models.MakeUser()

			user.Login = "Login"
			user.HashedPassword = "Password"
			ctx := context.Background()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			stor := Storage{driver: db}

			exp := mock.ExpectExec("INSERT INTO users (login, password) VALUES ($1, $2)").WithArgs(
				user.Login, user.HashedPassword,
			)
			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				exp.WillReturnResult(sqlmock.NewResult(1, 1))
			}

			err = stor.UserSet(ctx, user)
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

func TestUserGet(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
	}{
		{
			name:    "simple_happy_path",
			storErr: nil,
		},
		{
			name:    "when_stor_return_err",
			storErr: errors.New("some_err"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			login := "Login"
			ctx := context.Background()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			stor := Storage{driver: db}

			exp := mock.ExpectQuery("SELECT id, password FROM users WHERE login ilike $1 limit 1").WithArgs(
				login,
			)
			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				exp.WillReturnRows(sqlmock.NewRows([]string{"id", "password"}).AddRow("id", "pass"))
			}

			user, err := stor.UserGet(ctx, login)
			if ex.storErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, user.ID, "id")
				assert.Equal(t, user.HashedPassword, "pass")
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

		})

	}

}
