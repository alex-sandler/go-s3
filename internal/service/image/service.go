package image

import "go-s3/internal/config"

type Service struct {
	quality  int
	maxWidth int
}

func NewImageService(cfg *config.Config) *Service {
	return &Service{
		quality:  cfg.Image.Quality,
		maxWidth: cfg.Image.MaxWidth,
	}
}
