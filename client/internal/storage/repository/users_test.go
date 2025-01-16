package repository

import (
	"client/internal/models"
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
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
			user := models.User{ID: "123", Login: "Login", HashedPassword: "ququ"}
			ctx := context.Background()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			stor := Storage{driver: db}

			exp := mock.ExpectExec(userCreateSQL).WithArgs(user.ID, user.Login, user.HashedPassword)
			if ex.storErr != nil {
				exp.WillReturnError(ex.storErr)
			} else {
				exp.WillReturnResult(sqlmock.NewResult(1, 1))
			}

			err = stor.UserCreate(ctx, user)
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
				exp.WillReturnRows(sqlmock.NewRows([]string{"id", "login", "password"}).AddRow("id", "Login", "pass"))
			}

			user, err := stor.UserGet(ctx, login)
			if ex.storErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, user.ID, "id")
				assert.Equal(t, user.Login, login)
				assert.Equal(t, user.HashedPassword, "pass")
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

		})

	}

}
