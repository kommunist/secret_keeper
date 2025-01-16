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
			user := models.User{}

			user.Login = "Login"
			user.HashedPassword = "Password"
			ctx := context.Background()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			stor := Storage{driver: db}

			exp := mock.ExpectExec(userSetSQL).WithArgs(user.Login, user.HashedPassword)
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

			exp := mock.ExpectQuery(userGetSQL).WithArgs(login)
			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				exp.WillReturnRows(sqlmock.NewRows([]string{"id", "login", "password"}).AddRow("id", "login", "pass"))
			}

			user, err := stor.UserGet(ctx, login)
			if ex.storErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "id", user.ID)
				assert.Equal(t, "login", user.Login)
				assert.Equal(t, "pass", user.HashedPassword)
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

		})

	}

}
