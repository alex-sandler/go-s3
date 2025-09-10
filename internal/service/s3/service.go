package s3

import (
	"go-s3/internal/config"

	"github.com/minio/minio-go/v7"
)

type Service struct {
	client *minio.Client
	bucket string
}

func NewMinioService(client *minio.Client, cfg *config.Config) *Service {
	return &Service{
		client: client,
		bucket: cfg.Minio.BucketName,
	}
}
