package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/app/config"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/app/server"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	ctx := context.Background()

	cfg := config.LoadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	pool, dbConnectionErr := pgxpool.New(ctx, cfg.DB.Dsn)

	if dbConnectionErr != nil {
		slog.Error("database is failed to connect", "error", dbConnectionErr)
		os.Exit(1)
	}

	defer pool.Close()

	srv := server.New(cfg, pool)
	srvErr := srv.Run()

	if srvErr != nil {
		log.Fatal("Failed to start server: ", srvErr)
	}
}
