package secret

import (
	"client/internal/models"
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
		model   models.Secret
	}{
		{
			name:    "happy_path_list_secret",
			storErr: nil,
			model: models.Secret{
				ID:      "ID",
				Name:    "Name",
				Pass:    "Pass",
				Meta:    "Meta",
				Version: "Ver",
			},
		},
		{
			name:    "when_stor_return_error",
			storErr: errors.New("error"),
			model:   models.Secret{},
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			stor := NewMockSecretAccessor(gomock.NewController(t))

			item := Make(stor)

			userID := "userID"

			stor.EXPECT().SecretList(
				context.Background(), userID, "0",
			).Return([]models.Secret{ex.model}, ex.storErr)

			result, err := item.List(
				models.User{Login: "login", Password: "pass", ID: userID},
			)

			if ex.storErr == nil {
				assert.NoError(t, err)

				assert.Len(t, result, 1)
			} else {
				assert.Error(t, err)
			}
		})

	}

}
