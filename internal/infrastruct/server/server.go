package server

import (
	"context"
	"fmt"
	"go-s3/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {

	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.Port),
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
