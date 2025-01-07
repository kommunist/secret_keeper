package syncer

import (
	"client/internal/config"
	"client/internal/current"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestStartStop(t *testing.T) {
	t.Run("correct_start_and_stop", func(t *testing.T) {
		current.UnsetUser()

		c := config.Make()
		stor := NewMockStorageAccessor(gomock.NewController(t))
		roamer := NewMockRoamerAccessor(gomock.NewController(t))

		item := Make(c, stor, roamer)

		item.Start()
		time.Sleep(3 * time.Second)

		item.Stop()

	})

}
