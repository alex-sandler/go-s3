package service

import (
	"bytes"
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"

	"github.com/minio/minio-go/v7"
)

func (m *MinioService) UploadFile(ctx context.Context, fileName string, data []byte, contentType string) error {
	l := logger.FromContext(ctx)

	reader := bytes.NewReader(data)

	_, err := m.client.PutObject(ctx, m.bucket, fileName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		l.Errorf("service.UploadFile: failed to upload file: %v", err)
		return fmt.Errorf("service.UploadFile: failed to upload file: %w", err)
	}

	return nil
}
