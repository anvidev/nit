package main

import (
	"log/slog"
	"os"

	"github.com/anvidev/nit/config"
	"github.com/anvidev/nit/internal/application"
	"github.com/anvidev/nit/internal/store"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db := store.OpenDB(cfg)
	defer db.Close()

	app := application.New(cfg, logger, db)

	logger.Info("application starting", slog.String("port", cfg.Addr))

	if err := app.Run(); err != nil {
		logger.Error("error running application", slog.Any("error", err))
	}
}
