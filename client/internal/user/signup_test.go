package user

import (
	"client/internal/models"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUP(t *testing.T) {
	exList := []struct {
		name      string
		roamerErr error
	}{
		{
			name:      "simple_happy_path",
			roamerErr: nil,
		},
		{
			name:      "when_roamer_return_err",
			roamerErr: errors.New("ququ"),
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			stor := NewMockUserAccessor(gomock.NewController(t))
			roam := NewMockRemoteUserAccessor(gomock.NewController(t))

			user := models.User{Login: "123"}

			item := Make(stor, roam)

			roam.EXPECT().UserSet(user).Return(ex.roamerErr)

			result := item.SignUP(user)

			if ex.roamerErr == nil {
				assert.NoError(t, result)
			} else {
				assert.Error(t, result)
			}

		})
	}
}
