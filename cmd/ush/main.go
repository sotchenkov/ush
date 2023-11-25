package main

import (
	"fmt"
	"log/slog"
	"os"
	"ush/internal/config"
	"ush/internal/http-server/handlers/url/save"
	mwLogger "ush/internal/http-server/middleware/logger"
	"ush/internal/lib/logger/sl"
	"ush/internal/storage/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting ush", slog.String("env", cfg.Env))

	fmt.Println(cfg)

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init db", sl.Err(err))
		os.Exit(1)
	}

	// err = storage.SaveURL("https://www.google.com", "google")
	// if err != nil {
	// 	log.Error("Failed to init db", sl.Err(err))
	// 	os.Exit(1)
	// }

	_ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	// router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
