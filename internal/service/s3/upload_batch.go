package s3

import (
	"bytes"
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"

	"github.com/minio/minio-go/v7"
)

func (s *Service) UploadBatch(ctx context.Context, files map[string][]byte, contentType string) error {
	l := logger.FromContext(ctx)

	for fileName, data := range files {
		_, err := s.client.PutObject(ctx, s.bucket, fileName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			l.Errorf("service.UploadBatch: failed to upload file %s: %v", fileName, err)
			return fmt.Errorf("service.UploadBatch: failed to upload file %s: %w", fileName, err)
		}
	}
	return nil
}
