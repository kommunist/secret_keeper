package user

import (
	"client/internal/models"
	"context"
	"database/sql"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	exList := []struct {
		name          string
		localAuthErr  error
		roamerErr     error
		userCreateErr error
		realErr       bool
	}{
		{
			name:         "happy_path_when_locally",
			localAuthErr: nil,
			realErr:      false,
		},
		{
			name:         "when_stor_return_real_error",
			localAuthErr: errors.New("quququ"),
			realErr:      true,
		},
		{
			name:         "when_stor_return_no_rows_happy_path",
			localAuthErr: sql.ErrNoRows,
			roamerErr:    nil,
			realErr:      false,
		},
		{
			name:         "when_stor_return_no_rows_and_roamer_error",
			localAuthErr: sql.ErrNoRows,
			roamerErr:    errors.New("ququ"),
			realErr:      true,
		},
		{
			name:          "stor_return_no_rows_roamer_return_nil_and_create_return_err",
			localAuthErr:  sql.ErrNoRows,
			roamerErr:     nil,
			userCreateErr: errors.New("ququq"),
			realErr:       true,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			stor := NewMockUserAccessor(gomock.NewController(t))
			roam := NewMockRemoteUserAccessor(gomock.NewController(t))

			login := "123"
			password := "456"

			var expectedUser models.User

			item := Make(stor, roam, func(f models.User) { expectedUser = f })

			stor.EXPECT().UserGet(context.Background(), login).Return(
				models.User{Login: login, ID: "IDIDID"},
				ex.localAuthErr,
			)

			if ex.localAuthErr == sql.ErrNoRows {

				roaUser := models.User{
					Login: login, ID: "IDIDID", HashedPassword: "qwert",
				}

				roam.EXPECT().UserGet(models.User{Login: login, Password: password}).Return(
					roaUser, ex.roamerErr,
				)
				if ex.roamerErr == nil {
					stor.EXPECT().UserCreate(
						context.Background(), roaUser,
					).Return(ex.userCreateErr)
				}
			}

			err := item.SignIN(login, password)

			if ex.realErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, expectedUser.Login, login)
				assert.Equal(t, expectedUser.Password, password)
				assert.Equal(t, expectedUser.ID, "IDIDID")
				assert.NoError(t, err)
			}
		})
	}
}
