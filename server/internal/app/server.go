package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Основной метод пакета API
func (a *App) Start() error {
	slog.Info("server started", "URL", a.setting.ServerURL)

	a.listenInterrupt()

	// err := a.Server.ListenAndServeTLS(a.setting.CertPath, a.setting.CertKeyPath)
	err := a.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start with tls", "err", err)
		return err
	}
	return nil
}

// Функция, останавливающая сервер
func (a *App) stop() {
	err := a.Server.Shutdown(context.Background())
	if err != nil {
		slog.Error("Error when shutdown server", "err", err)
	}
}

// Механизм прослушивания прерываний
func (i *App) listenInterrupt() {
	sigint := make(chan os.Signal, 3)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go i.waitInterrupt(sigint)
}

func (i *App) waitInterrupt(sigint chan os.Signal) {
	<-sigint

	i.stop()

	close(sigint)
}
