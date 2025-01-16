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

// --- Mock для генерации версии
type MockVer struct{}

func (i *MockVer) Get() string { return "version" }

// --- Mock для генерации версии

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
			item.verGet = &MockVer{}

			secret := models.Secret{}
			secret.ID = "ID"
			secret.Meta = "meta"
			secret.Name = "name"
			secret.Pass = "pass"
			secret.UserID = "userID"

			expSecret := secret
			expSecret.UserID = "userID"
			expSecret.Version = "version"

			current.SetUser(models.User{Login: "login", Password: "password", ID: "userID"})

			defer current.UnsetUser()

			stor.EXPECT().SecretsUpsert(context.Background(), []models.Secret{expSecret}).
				Return(ex.storErr)

			err := item.Upsert(secret)

			if ex.storErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})

	}

}
