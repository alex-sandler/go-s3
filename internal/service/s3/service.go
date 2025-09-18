package s3

import (
	"go-s3/internal/config"
	"go-s3/internal/service"

	"github.com/minio/minio-go/v7"
)

type Service struct {
	client       *minio.Client
	bucket       string
	imageService service.ImageService
}

func NewMinioService(client *minio.Client, cfg *config.Config, imageService service.ImageService) *Service {
	return &Service{
		client:       client,
		bucket:       cfg.Minio.BucketName,
		imageService: imageService,
	}
}
