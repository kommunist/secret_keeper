package userset

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/internal/models"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	exList := []struct {
		name      string
		storTimes int
		storErr   error
		hashErr   error
		status    int
	}{
		{
			name:      "create_user_happy_path",
			storTimes: 1,
			status:    http.StatusOK,
		},
		{
			name:      "when_hashed_pass_err",
			storTimes: 0,
			hashErr:   errors.New("ququ"),
			status:    http.StatusInternalServerError,
		},
		{
			name:      "when_stor_err",
			storTimes: 1,
			storErr:   errors.New("ququ"),
			status:    http.StatusInternalServerError,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockUserSetter(ctrl)
			hasher := NewMockPasswordHasher(ctrl)

			h := Make(stor)
			h.hasher = hasher

			hasher.EXPECT().HashPassword("password").Return("hashedPass", ex.hashErr)

			ctx := context.Background()

			user := models.User{Login: "login", Password: "password"}
			input, _ := json.Marshal(user)

			request := httptest.NewRequest(http.MethodPost, "/any", bytes.NewReader(input)).WithContext(ctx)

			user.HashedPassword = "hashedPass"
			stor.EXPECT().UserSet(ctx, user).Times(ex.storTimes).Return(ex.storErr)

			w := httptest.NewRecorder()
			h.Handler(w, request)
			res := w.Result()

			assert.Equal(t, ex.status, res.StatusCode)
		})

	}
}
