package secretset

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/internal/auth"
	"server/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	exList := []struct {
		name              string
		inputJSON         []byte
		upsertReceive     []models.Secret
		storTimes         int
		storErr           error
		status            int
		withUserInContext bool
	}{
		{
			name: "create_secrets_happy_path",
			inputJSON: func() []byte {
				data, _ := json.Marshal(
					[]models.Secret{{ID: "id", Name: "name", Pass: "pass", Meta: "meta", Version: "ver"}},
				)
				return data
			}(),
			upsertReceive: []models.Secret{
				{ID: "id", Name: "name", Pass: "pass", Meta: "meta", Version: "ver", UserID: "user_id"},
			},
			storTimes:         1,
			status:            http.StatusOK,
			withUserInContext: true,
		},
		{
			name: "stor_return_err",
			inputJSON: func() []byte {
				data, _ := json.Marshal(
					[]models.Secret{{ID: "id", Name: "name", Pass: "pass", Meta: "meta", Version: "ver"}},
				)
				return data
			}(),
			upsertReceive: []models.Secret{
				{ID: "id", Name: "name", Pass: "pass", Meta: "meta", Version: "ver", UserID: "user_id"},
			},
			storTimes:         1,
			storErr:           errors.New("ququ"),
			status:            http.StatusInternalServerError,
			withUserInContext: true,
		},
		{
			name: "stor_return_err",
			inputJSON: func() []byte {
				data, _ := json.Marshal(
					[]models.Secret{{ID: "id", Name: "name", Pass: "pass", Meta: "meta", Version: "ver"}},
				)
				return data
			}(),
			upsertReceive: []models.Secret{
				{ID: "id", Name: "name", Pass: "pass", Meta: "meta", Version: "ver", UserID: "user_id"},
			},
			storTimes:         0,
			status:            http.StatusInternalServerError,
			withUserInContext: false,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockSecretUpserter(ctrl)

			h := Make(stor)

			ctx := context.Background()
			if ex.withUserInContext {
				ctx = context.WithValue(ctx, auth.UserIDKey, models.User{Login: "Login", ID: "user_id"})
			}

			request := httptest.NewRequest(
				http.MethodPost, "/any", bytes.NewReader(ex.inputJSON),
			).WithContext(ctx)

			stor.EXPECT().SecretUpsert(ctx, ex.upsertReceive).Times(ex.storTimes).Return(ex.storErr)

			w := httptest.NewRecorder()
			h.Handler(w, request)
			res := w.Result()

			assert.Equal(t, ex.status, res.StatusCode)
		})

	}
}
