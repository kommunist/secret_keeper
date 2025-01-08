package app

import (
	"net/http"
	"server/internal/config"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
)

func TestStartStop(t *testing.T) {
	t.Run("start_stop_check", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := NewMockStorager(ctrl)

		h := App{
			setting: config.MainConfig{},
			storage: stor,
		}
		h.Server = http.Server{Addr: "localhost:1024", Handler: h.initRouter()}

		go func() { time.Sleep(3 * time.Second); h.stop() }()

		h.Start()

	})

}
