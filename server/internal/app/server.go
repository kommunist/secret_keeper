package app

import (
	"context"
	"log/slog"
	"net/http"
)

// Основной метод пакета API
func (a *App) Start() error {
	slog.Info("server started", "URL", a.setting.ServerURL)

	err := a.Server.ListenAndServeTLS(a.setting.CertPath, a.setting.CertKeyPath)
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start with tls", "err", err)
		return err
	}
	return nil
}

// Функция, останавливающая сервер
func (a *App) Stop() {
	err := a.Server.Shutdown(context.Background())
	if err != nil {
		slog.Error("Error when shutdown server", "err", err)
	}
}
