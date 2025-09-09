package grpc

import (
	"context"
	"fmt"
	"go-s3/internal/config"
	"go-s3/internal/infrastruct/logger"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	cfg    *config.Config
}

func NewServer(cfg *config.Config) *Server {
	server := grpc.NewServer()

	return &Server{
		server: server,
		cfg:    cfg,
	}
}

func (s *Server) Run(ctx context.Context) error {
	l := logger.FromContext(ctx)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.cfg.GRPC.Host, s.cfg.GRPC.Port))
	if err != nil {
		l.Errorf("server.Run: failed to listen: %v", err)
		return fmt.Errorf("server.Run: failed to listen: %w", err)
	}

	go func() {
		<-ctx.Done()
		s.server.GracefulStop()
	}()

	l.Info("Server is running")

	if err := s.server.Serve(listener); err != nil {
		l.Errorf("server.Run: failed to serve: %v", err)
		return fmt.Errorf("server.Run: failed to serve: %w", err)
	}

	l.Info("Server is stopped gracefully")

	return nil
}
