package s3

import (
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"

	"github.com/minio/minio-go/v7"
)

func (s *Service) ListFiles(ctx context.Context) ([]string, error) {
	var files []string
	l := logger.FromContext(ctx)

	objectCh := s.client.ListObjects(ctx, s.bucket, minio.ListObjectsOptions{})

	for object := range objectCh {
		if object.Err != nil {
			l.Errorf("service.ListFiles: failed to list objects: %v", object.Err)
			return nil, fmt.Errorf("service.ListFiles: failed to list objects: %w", object.Err)
		}
		files = append(files, object.Key)
	}

	return files, nil
}
