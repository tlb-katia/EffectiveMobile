package main

import (
	"EffectiveMobile_Project/cmd/migrator"
	"EffectiveMobile_Project/config"
	"EffectiveMobile_Project/internal/server"
	"EffectiveMobile_Project/pkg/repository"
	"EffectiveMobile_Project/pkg/storage/postgres"
	"github.com/go-chi/chi/v5"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"log/slog"
	"os"
)

// @title Effective Mobile Project
// @version 1.0
// @description This is a simple REST Api to manage a list of songs

// @contact.name Katja K
// @contact.url https://github.com/tlb-katia
// @contact.email tlb-kei7@yandex.ru

// @host localhost:8080
// @BasePath /

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger("local")
	log.Info("starting server", slog.String("address", cfg.Address))

	psql, err := postgres.NewStorage(cfg, log)
	if err != nil {
		log.Error("failed to init psql", err)
		os.Exit(1)
	}
	defer psql.Db.Close()

	migrator.RunMigrations(cfg.PGName, psql.Driver, cfg.MigrationPath)

	repo := repository.NewRepository(psql)
	srv := server.NewServer(repo, chi.NewRouter(), log)
	srv.Run(cfg)

	log.Info("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
