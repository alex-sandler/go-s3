package s3

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (s *Service) DeleteFile(ctx context.Context, fileName string) error {
	return s.client.RemoveObject(ctx, s.bucket, fileName, minio.RemoveObjectOptions{})
}
