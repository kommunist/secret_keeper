package user

import (
	"client/internal/current"
	"client/internal/encrypt"
	"client/internal/models"
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	f := models.MakeUser()
	f.Login = "Login"
	f.Password = "Password"

	exList := []struct {
		name         string
		form         models.User
		returnedID   string
		retrunedPass string
		returnedErr  error
		wihError     bool
	}{
		{
			name:         "happy_path_correct_set",
			form:         f,
			returnedID:   "ID",
			retrunedPass: func() string { str, _ := encrypt.HashPassword("Password"); return str }(),
			returnedErr:  nil,
			wihError:     false,
		},
		{
			name:         "incorrect_password",
			form:         f,
			returnedID:   "ID",
			retrunedPass: func() string { str, _ := encrypt.HashPassword("Ququq"); return str }(),
			returnedErr:  nil,
			wihError:     true,
		},
		{
			name:         "stor_returned_err",
			form:         f,
			returnedID:   "",
			retrunedPass: "",
			returnedErr:  errors.New("stor_error"),
			wihError:     true,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			defer current.UnsetUser()

			stor := NewMockUserAccessor(gomock.NewController(t))
			item := Make(stor)

			stor.EXPECT().UserGet(context.Background(), "Login").Return(ex.returnedID, ex.retrunedPass, ex.returnedErr)

			err := item.SignIN(ex.form)

			if ex.wihError {
				assert.Error(t, err)
				assert.NotEqual(t, "Login", current.User.Login)
				assert.NotEqual(t, "Password", current.User.Password)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "Login", current.User.Login)
				assert.Equal(t, "Password", current.User.Password)
			}
		})

	}

}
