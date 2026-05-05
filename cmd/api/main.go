package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/volodymyr-miretskyi/go-nutrition/internal/app/config"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	cfg := config.LoadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, dbConnectionErr := pgx.Connect(ctx, cfg.DB.Dsn)

	if dbConnectionErr != nil {
		slog.Error("database is failed to connect", "error", dbConnectionErr)
		os.Exit(1)
	}

	defer conn.Close(ctx)
}
