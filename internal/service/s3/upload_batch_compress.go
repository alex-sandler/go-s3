package s3

import (
	"bytes"
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"

	"github.com/minio/minio-go/v7"
)

func (s *Service) UploadBatchCompress(ctx context.Context, files map[string][]byte, contentType string) error {
	l := logger.FromContext(ctx)

	compressedFiles, err := s.imageService.CompressBatch(ctx, files)
	if err != nil {
		l.Errorf("service.UploadBatchCompress: failed to compress files: %v", err)
		return fmt.Errorf("service.UploadBatchCompress: failed to compress files: %w", err)
	}

	for fileName, data := range compressedFiles {
		_, err := s.client.PutObject(ctx, s.bucket, fileName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			l.Errorf("service.UploadBatchCompress: failed to upload file %s: %v", fileName, err)
			return fmt.Errorf("service.UploadBatchCompress: failed to upload file %s: %w", fileName, err)
		}
	}
	return nil
}
