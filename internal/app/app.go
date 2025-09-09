package app

import (
	"context"
	"fmt"
	"go-s3/internal/config"
	"go-s3/internal/infrastruct/grpc"
	"go-s3/internal/infrastruct/logger"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

var defaultLevel = zap.NewAtomicLevelAt(zap.InfoLevel)

func Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	l := logger.New(defaultLevel, os.Stdout)
	ctx = logger.WithLogger(ctx, l)

	cfg, err := config.Load()
	if err != nil {
		l.Errorf("app.Run: failed to load config: %v", err)
		return fmt.Errorf("app.Run: %w", err)
	}

	grpcServer := grpc.NewServer(cfg)

	if err := grpcServer.Run(ctx); err != nil {
		l.Errorf("app.Run: failed to run server: %v", err)
		return fmt.Errorf("app.Run: %w", err)
	}

	return nil
}
