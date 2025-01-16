package syncer

import (
	"client/internal/config"
	"client/internal/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestStartStop(t *testing.T) {
	t.Run("correct_start_and_stop", func(t *testing.T) {
		c := config.Make()
		stor := NewMockStorageAccessor(gomock.NewController(t))
		roamer := NewMockRoamerAccessor(gomock.NewController(t))
		fcu := func() models.User { return models.User{} }

		item := Make(c, stor, roamer, fcu)

		item.Start()
		time.Sleep(3 * time.Second)

		item.Stop()

	})

}
