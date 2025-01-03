package secret

import (
	"client/internal/current"
	"client/internal/models"
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpsert(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
	}{
		{
			name:    "happy_path_upsert_secret",
			storErr: nil,
		},
		{
			name:    "when_stor_return_error",
			storErr: errors.New("error"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			stor := NewMockSecretAccessor(gomock.NewController(t))

			item := Make(stor)

			form := models.MakeSecret()
			form.ID = "id"
			form.Meta = "meta"
			form.Name = "name"
			form.Pass = "pass"

			current.SetUser("login", "password", "userID")
			defer current.UnsetUser()

			stor.EXPECT().SecretUpsert(
				context.Background(),
				form.ID, form.Name, form.Pass, form.Meta, current.User.ID, gomock.Any(),
			).Return(ex.storErr)

			err := item.Upsert(form)

			if ex.storErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})

	}

}
