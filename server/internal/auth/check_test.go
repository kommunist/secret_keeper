package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/internal/encrypt"
	"server/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func testHandler(w http.ResponseWriter, r *http.Request) {}

func TestCheck(t *testing.T) {
	exList := []struct {
		name       string
		hashedPass string
		storErr    error
		resStatus  int
	}{
		{
			name:       "happy_path",
			hashedPass: func() string { res, _ := encrypt.Item{}.HashPassword("Password"); return res }(),
			resStatus:  http.StatusOK,
		},
		{
			name:       "when_stor_return_err",
			storErr:    errors.New("ququ"),
			hashedPass: func() string { res, _ := encrypt.Item{}.HashPassword("Password"); return res }(),
			resStatus:  http.StatusInternalServerError,
		},
		{
			name:       "when_incorrect_pass",
			hashedPass: "qwe",
			resStatus:  http.StatusUnauthorized,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctx := context.Background()
			stor := NewMockUserGetter(gomock.NewController(t))
			item := Make(stor)

			nextHandler := http.HandlerFunc(testHandler)
			handlerToTest := item.Check(nextHandler)

			model := models.User{HashedPassword: ex.hashedPass}

			stor.EXPECT().UserGet(ctx, "LoginOfUser").Return(model, ex.storErr)

			request := httptest.NewRequest(http.MethodGet, "/any", nil).WithContext(ctx)
			request.Header.Set("Login", "LoginOfUser")
			request.Header.Set("Password", "Password")

			w := httptest.NewRecorder()
			handlerToTest.ServeHTTP(w, request)

			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, ex.resStatus, res.StatusCode)

		})

	}

}
