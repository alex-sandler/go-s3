package service

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (m *MinioService) DeleteFile(ctx context.Context, fileName string) error {
	return m.client.RemoveObject(ctx, m.bucket, fileName, minio.RemoveObjectOptions{})
}
