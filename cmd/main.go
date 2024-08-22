package main

import (
	"log/slog"
	"os"

	"github.com/anvidev/nit/config"
	"github.com/anvidev/nit/internal/application"
	"github.com/anvidev/nit/internal/store"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Warn("error loading .env file", slog.Any("error", err))
	}

	db, err := store.OpenDB(cfg)
	if err != nil {
		logger.Warn("error connecting to db", slog.Any("error", err))
	}
	defer db.Close()

	s3, err := store.NewAwsS3Bucket(cfg)
	if err != nil {
		logger.Error("error initializing s3 client", slog.Any("error", err))
		os.Exit(1)
	}

	app := application.New(cfg, logger, db, s3)

	logger.Info("application starting", slog.String("port", cfg.Addr))

	if err := app.Run(); err != nil {
		logger.Error("error running application", slog.Any("error", err))
	}
}
