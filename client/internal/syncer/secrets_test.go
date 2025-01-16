package syncer

import (
	"client/internal/config"
	"client/internal/current"
	"client/internal/models"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

// --- Mock для генерации версии
type MockVer struct{}

func (i *MockVer) Get() string { return "version" }

func TestSyncSecrets(t *testing.T) {
	exList := []struct {
		name               string
		userSeted          bool
		lastSyncError      error
		secretListErr      error
		roamerSecretSetErr error
		roamerSecretGetErr error
		secretsUpsertErr   error
		saveSyncEventErr   error
	}{
		{
			name:      "when_user_not_seted",
			userSeted: false,
		},
		{
			name:      "when_user_seted_and_happy_path",
			userSeted: true,
		},
		{
			name:          "when_user_seted_and_last_sync_return_err",
			userSeted:     true,
			lastSyncError: errors.New("qq"),
		},
		{
			name:          "when_user_seted_and_secret_list_return_err",
			userSeted:     true,
			secretListErr: errors.New("qq"),
		},
		{
			name:               "when_user_seted_and_roamer_secret_set_return_err",
			userSeted:          true,
			roamerSecretSetErr: errors.New("qq"),
		},
		{
			name:               "when_user_seted_and_roamer_secret_get_return_err",
			userSeted:          true,
			roamerSecretGetErr: errors.New("qq"),
		},
		{
			name:             "when_user_seted_and_secrets_upsert_return_err",
			userSeted:        true,
			secretsUpsertErr: errors.New("qq"),
		},
		{
			name:             "when_user_seted_and_save_sync_event_return_err",
			userSeted:        true,
			saveSyncEventErr: errors.New("qq"),
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			defer current.UnsetUser()

			if ex.userSeted {
				cu := models.User{ID: "123", Login: "login", Password: "pass"}
				current.SetUser(cu)
			}

			c := config.MainConfig{}

			stor := NewMockStorageAccessor(gomock.NewController(t))
			roamer := NewMockRoamerAccessor(gomock.NewController(t))

			item := Make(&c, stor, roamer)
			item.verGet = &MockVer{}

			if !ex.userSeted {
				item.syncSecrets()
				return
			}

			stor.EXPECT().GetLastSyncEventVersion(context.Background(), "secret").Return("0", ex.lastSyncError)

			if ex.lastSyncError != nil {
				item.syncSecrets()
				return
			}

			// --sendLocalSecrets--

			secrets := []models.Secret{{Name: "name", Pass: "pass", Meta: "meta", Version: "version"}}
			stor.EXPECT().SecretList(context.Background(), current.User.ID, "0").Return(secrets, ex.secretListErr)

			if ex.secretListErr != nil {
				item.syncSecrets()
				return
			}

			roamer.EXPECT().SecretSet(secrets).Return(ex.roamerSecretSetErr)
			if ex.roamerSecretSetErr != nil {
				item.syncSecrets()
				return
			}

			// --sendLocalSecrets--

			// --getServerSecrets--

			secrets = []models.Secret{{Name: "name", Pass: "pass", Meta: "meta", Version: "version"}}
			roamer.EXPECT().SecretGet("0").Return(secrets, ex.roamerSecretGetErr)
			if ex.roamerSecretGetErr != nil {
				item.syncSecrets()
				return
			}

			stor.EXPECT().SecretsUpsert(context.Background(), secrets).Return(ex.secretsUpsertErr)
			if ex.secretsUpsertErr != nil {
				item.syncSecrets()
				return
			}

			// --sendLocalSecrets--

			stor.EXPECT().SaveSyncEvent(context.Background(), "secret", "version").Return(ex.saveSyncEventErr)
			if ex.saveSyncEventErr != nil {
				item.syncSecrets()
				return
			}

			item.syncSecrets()
		})
	}

}
