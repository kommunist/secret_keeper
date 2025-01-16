package secretget

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/internal/auth"
	"server/internal/models"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	exList := []struct {
		name              string
		withUserInContext bool
		storTimes         int
		storErr           error
		resStatus         int
		resSecrets        []models.Secret
	}{
		{
			name:              "simple_happy_path",
			withUserInContext: true,
			storTimes:         1,
			resStatus:         http.StatusOK,
			resSecrets: []models.Secret{
				{
					ID:      "id",
					Name:    "name",
					Pass:    "pass",
					Version: "ver",
					Meta:    "meta",
				},
			},
		},
		{
			name:              "without_user_in_context",
			withUserInContext: false,
			storTimes:         0,
			resStatus:         http.StatusInternalServerError,
		},
		{
			name:              "when_stor_return_err",
			withUserInContext: true,
			storTimes:         1,
			storErr:           errors.New("ququq"),
			resStatus:         http.StatusInternalServerError,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockSecretGetter(ctrl)

			h := Make(stor)

			ctx := context.Background()
			if ex.withUserInContext {
				ctx = context.WithValue(ctx, auth.UserIDKey, models.User{Login: "Login", ID: "user_id"})
			}

			request := httptest.NewRequest(http.MethodGet, "/any?version=0", nil).WithContext(ctx)

			stor.EXPECT().SecretGet(ctx, "user_id", "0").Times(ex.storTimes).Return(
				ex.resSecrets, ex.storErr,
			)

			w := httptest.NewRecorder()
			h.Handler(w, request)
			res := w.Result()

			assert.Equal(t, ex.resStatus, res.StatusCode)
		})

	}

}
