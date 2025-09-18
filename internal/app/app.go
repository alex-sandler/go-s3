package app

import (
	"context"
	"errors"
	"fmt"
	"go-s3/internal/config"
	"go-s3/internal/controller"
	"go-s3/internal/infrastruct/logger"
	"go-s3/internal/infrastruct/minio"
	"go-s3/internal/infrastruct/server"
	"go-s3/internal/service/image"
	"go-s3/internal/service/s3"
	"net/http"
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

	minioClient := minio.NewMinioClient()
	err = minioClient.Init(ctx, cfg)
	if err != nil {
		l.Errorf("app.Run: failed to init minio client: %v", err)
		return fmt.Errorf("app.Run: %w", err)
	}
	imageService := image.NewImageService(cfg)
	s3Service := s3.NewMinioService(minioClient.GetClient(), cfg, imageService)
	c := controller.NewController(s3Service, imageService)

	srv := server.NewServer(cfg, c)
	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			l.Errorf("app.Run: failed to start server server: %v", err)
		}
	}()
	l.Info("Server started, waiting for shutdown signal...")

	<-ctx.Done()

	l.Info("Shutdown signal received, stopping server...")
	return srv.Stop(ctx)
}
