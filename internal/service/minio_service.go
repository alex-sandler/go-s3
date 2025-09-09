package service

import (
	"github.com/minio/minio-go/v7"
)

type MinioService struct {
	client *minio.Client
	bucket string
}

func NewMinioService(client *minio.Client, bucket string) *MinioService {
	return &MinioService{
		client: client,
		bucket: bucket,
	}
}
