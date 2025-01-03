package user

import (
	"client/internal/models"
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUP(t *testing.T) {
	form := models.MakeUser()
	form.Login = "Login"
	form.Password = "Password"

	exList := []struct {
		name        string
		returnedErr error
	}{
		{
			name:        "simple_happy_path",
			returnedErr: nil,
		},
		{
			name:        "when_error_from_storage",
			returnedErr: errors.New("new_error"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			stor := NewMockUserAccessor(gomock.NewController(t))
			item := Make(stor)

			// Пока сделал gomock.Any потому что не знаю, как замокать вызов хеширующей функции
			stor.EXPECT().UserCreate(context.Background(), "Login", gomock.Any()).Return(ex.returnedErr)

			err := item.SignUP(form)
			if ex.returnedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}
