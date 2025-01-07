package syncer

import (
	"client/internal/logger"
	"time"
)

func (i *Item) Start() {
	logger.Logger.Info("Syncer started")

	ticker := time.NewTicker(2 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				logger.Logger.Info("Syncer tick!")
				i.syncSecrets()
			case <-i.stoper:
				ticker.Stop()
				return
			}
		}
	}()
}

func (i *Item) Stop() {
	logger.Logger.Info("Stop syncer")

	i.stoper <- true
}
