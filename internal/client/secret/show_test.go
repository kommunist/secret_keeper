package secret

import (
	"context"
	"errors"
	"secret_keeper/internal/client/models"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	exList := []struct {
		name    string
		storErr error
		model   models.Secret
	}{
		{
			name:    "happy_path_show_secret",
			storErr: nil,
			model: models.Secret{
				ID:   "ID",
				Name: "Name",
				Pass: "Pass",
				Meta: "Meta",
				Ver:  "Ver",
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

			secretID := "secretID"

			stor.EXPECT().SecretShow(
				context.Background(), secretID,
			).Return(ex.model, ex.storErr)

			f, err := item.Show(secretID)

			if ex.storErr == nil {
				assert.NoError(t, err)

				assert.Equal(t, f.ID, ex.model.ID)
				assert.Equal(t, f.Name, ex.model.Name)
				assert.Equal(t, f.Pass, ex.model.Pass)
				assert.Equal(t, f.Meta, ex.model.Meta)
				assert.Equal(t, f.Ver, ex.model.Ver)

			} else {
				assert.Error(t, err)
			}
		})

	}

}
