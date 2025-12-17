package main

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/dimplesY/goose_test/internal/env"
	"github.com/dimplesY/goose_test/internal/helper"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	ctx := context.Background()

	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,   // 单文件最大体积（MB）
		MaxBackups: 5,    // 保留的旧文件数量
		MaxAge:     30,   // 文件保留天数
		Compress:   true, // 压缩旧日志
	}

	defer lumberjackLogger.Close()

	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)

	logger := slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	helper.InitHelper()

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
