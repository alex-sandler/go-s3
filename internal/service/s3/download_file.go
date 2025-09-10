package s3

import (
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"
	"io"

	"github.com/minio/minio-go/v7"
)

func (m *Service) DownloadFile(ctx context.Context, fileName string) ([]byte, error) {
	l := logger.FromContext(ctx)

	object, err := m.client.GetObject(ctx, m.bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		l.Errorf("service.DownloadFile: failed to download file: %v", err)
		return nil, fmt.Errorf("service.DownloadFile: failed to download file: %w", err)
	}
	defer object.Close()

	return io.ReadAll(object)
}
