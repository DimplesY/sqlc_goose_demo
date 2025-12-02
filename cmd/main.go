package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/dimplesY/goose_test/internal/env"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(ctx, env.GetEnvByName("GOOSE_DBSTRING", "postgres://postgres:123456@localhost:5432/goose_test"))

	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	logger.Info("database connected")

	api := application{db: conn}

	if err := api.run(api.mount()); err != nil {

		slog.Error("server failed to start", "error", err)

		os.Exit(1)
	}

}
