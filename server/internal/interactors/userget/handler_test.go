package userget

import (
	"context"
	"net/http"
	"net/http/httptest"
	"server/internal/auth"
	"server/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	exList := []struct {
		name              string
		withUserInContext bool
		storErr           error
		resStatus         int
	}{
		{
			name:              "simple_happy_path",
			withUserInContext: true,
			resStatus:         http.StatusOK,
		},
		{
			name:              "without_user_in_context",
			withUserInContext: false,
			resStatus:         http.StatusInternalServerError,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			h := Make()

			ctx := context.Background()
			if ex.withUserInContext {
				ctx = context.WithValue(ctx, auth.UserIDKey, models.User{Login: "Login", ID: "user_id"})
			}

			request := httptest.NewRequest(http.MethodGet, "/any", nil).WithContext(ctx)

			w := httptest.NewRecorder()
			h.Handler(w, request)
			res := w.Result()

			assert.Equal(t, ex.resStatus, res.StatusCode)
		})

	}

}
