package app

import (
	"client/internal/logger"
	"os"
	"os/signal"
	"syscall"
)

// Запуск приложения
func (a *App) Start() error {
	a.listenInterrupt()

	a.syncer.Start()

	a.tui.Hello()
	err := a.tui.Start()
	if err != nil {
		logger.Logger.Error("Start: when start tui", "err", err)
		return err
	}
	return nil
}

// Остановка приложения
func (a *App) stop() {
	a.syncer.Stop()
	a.tui.Stop()
}

// Обработчик прерываний
func (a *App) listenInterrupt() {
	sigint := make(chan os.Signal, 3)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go a.waitInterrupt(sigint)
}

// Слушатель прерываний
func (a *App) waitInterrupt(sigint chan os.Signal) {
	<-sigint
	a.stop()
	close(sigint)
}
