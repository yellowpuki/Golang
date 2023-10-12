package main

import (
	"os"

	"github.com/yellowpuki/simple-pg/internal/config"
	"golang.org/x/exp/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	// TODO: Config
	cfg := config.MustLoad()

	// TODO: Logger
	log := setupLogger(cfg.Env)
	log.Info("Starting application.", slog.String("env", cfg.Env))

	// TODO: Storage
	// TODO: Router
	// TODO: Server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
