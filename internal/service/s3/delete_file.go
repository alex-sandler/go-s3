package s3

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (m *Service) DeleteFile(ctx context.Context, fileName string) error {
	return m.client.RemoveObject(ctx, m.bucket, fileName, minio.RemoveObjectOptions{})
}
