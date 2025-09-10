package server

import (
	"context"
	"fmt"
	"go-s3/internal/config"
	"go-s3/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
}

func NewServer(cfg *config.Config, ctrl *controller.Controller) *Server {
	app := fiber.New(fiber.Config{
		AppName: "S3 Service",
	})

	ctrl.Routes(app)

	return &Server{
		app: app,
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf("%s:%s", s.cfg.HTTP.Host, s.cfg.HTTP.Port))
}

func (s *Server) Stop(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
